/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package agent

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/carbonaut/pkg/agent/targets/aws"
	"github.com/carbonaut/pkg/agent/targets/azure"
	"github.com/carbonaut/pkg/agent/targets/carbonaware"
	"github.com/carbonaut/pkg/agent/targets/gcp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type CarbonautAgentTarget interface {
	Register(any, *prometheus.Registry) error
	UnmarshalSpec([]byte) (any, error)
	GetTargetType() string
}

// interimTargetConfig used to load unrecognized config files before its clear which type is defined
type transportTargetConfig struct {
	// TODO: version which is used; CURRENTLY NOT USED (remove?)
	Version string `yaml:"version"`
	// type of target; its possible to configure multiple targets of the same type in one agent
	Type string `yaml:"type"`
	// identifying name which needs to be unique
	Name TargetName `yaml:"name"`
	// port thats used to expose target metrics
	Port int `yaml:"port"`
	// target spec
	Spec interface{} `yaml:"spec"`
	// target spec in bytes
	SpecBytes []byte
}

// targetProcess defines a current running process of a target
type targetProcess struct {
	interruptChan chan int
	status        targetProcessStatus
	cfg           transportTargetConfig
}

type targetProcessStatus string

var (
	// process is stateRunning
	stateRunning targetProcessStatus = "RUNNING"
	// process is currently stateStopped and data gets not pulled
	stateStopped targetProcessStatus = "STOPPED"
)

// name of the target used for identification
type TargetName string

type action string

const (
	// target should start
	ActionStart action = "START_TARGET"
	// target should stop
	ActionStop action = "STOP_TARGET"
	// all targets should stop; agent shuts down
	ActionStopAgent action = "STOP_AGENT_ALL_TARGETS"
	// load configuration file again (reload targets)
	ActionLoadConfig action = "LOAD_CONFIGURATION_FILE"
	// (internal only) target threw an error during execution
	actionError action = "ERROR_IN_TARGET_PROCESS"
)

type CommsChannel struct {
	// Action to perform
	Action action
	// Name of the target
	Name TargetName
	// further Details about the action for traceability
	Details string
	// optional field to specify an error
	err error
}

type promServeConfig struct {
	ServerReadTimeout  time.Duration
	ServerWriteTimeout time.Duration
	Port               int
	Registry           *prometheus.Registry
}

var (
	filename                  = "config.yaml"
	separator                 = "---"
	targetErrorRestartTimeout = 5 * time.Second
	targetImplementations     = []CarbonautAgentTarget{carbonaware.Target{}, aws.Target{}, azure.Target{}, gcp.Target{}}
	promServerReadTimeout     = 10 * time.Second
	promServerWriteTimeout    = 10 * time.Second
)

func Run(commsChan chan CommsChannel) error {
	log.Info().Msg("Starting Agent targets...")

	// blocks this process at the end and only gets "lifted" if action: ActionStopAgent gets called
	stopAgent := make(chan int)

	// start listening the comms channel for actions to take
	go commsRelay(commsChan, stopAgent)

	// initial start of configured targets
	commsChan <- CommsChannel{
		Action:  ActionLoadConfig,
		Details: "initial config loading",
	}

	// block main process until stop agent gets called
	<-stopAgent
	return nil
}

// commsRelay used to receive general tasks to execute and further distribute the work
func commsRelay(commsChan chan CommsChannel, stopAgent chan int) {
	// list of running targets that can be stopped by sending something through the channel
	currentTargetProcesses := map[TargetName]*targetProcess{}

	// listening the comms channel for actions to take
	for i := range commsChan {
		log.Info().Msg(i.Details)
		switch i.Action {
		case ActionStart:
			// action: start a task; the task needs to be configured
			p, err := startTarget(commsChan, &currentTargetProcesses[i.Name].cfg)
			if err != nil {
				log.Error().Err(err).Msg("unable to start target")
			}
			currentTargetProcesses[i.Name] = p

		case actionError:
			// action: task threw an error
			log.Error().Err(i.err).Msgf("timeout to restart target set to %d", targetErrorRestartTimeout)
			// TODO: this blocks the main thread...
			time.Sleep(targetErrorRestartTimeout)
			p, err := startTarget(commsChan, &currentTargetProcesses[i.Name].cfg)
			if err != nil {
				log.Error().Err(err).Msg("unable to start target after error")
			}
			currentTargetProcesses[i.Name] = p

		case ActionStop:
			// action: stop running task
			currentTargetProcesses = stopTarget(currentTargetProcesses, i.Name)

		case ActionStopAgent:
			for range currentTargetProcesses {
				currentTargetProcesses = stopTarget(currentTargetProcesses, i.Name)
			}
			stopAgent <- 1

		case ActionLoadConfig:
			// action: load the configuration file and adjust tasks according to the configuration
			updatedTargetProcessList, err := reloadConfig(commsChan, currentTargetProcesses)
			if err != nil {
				log.Fatal().Err(err).Msg("unable to reload provided configuration")
			}
			currentTargetProcesses = updatedTargetProcessList

		default:
			log.Error().Msgf("unrecognized action request received %s", string(i.Action))
		}
	}
}

// reloadConfig pulls current configuration file and runs it
func reloadConfig(commsChan chan CommsChannel, currentTargetProcesses map[TargetName]*targetProcess) (map[TargetName]*targetProcess, error) {
	cfg, err := loadConfiguration(filename, separator)
	if err != nil {
		return nil, err
	}
	newTargetProcesses, err := execTargetConfig(commsChan, currentTargetProcesses, cfg)
	if err != nil {
		return nil, err
	}
	return newTargetProcesses, nil
}

func loadConfiguration(filename, separator string) (map[TargetName]*transportTargetConfig, error) {
	// read configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// split configuration files for agents
	interimConfigs := map[TargetName]*transportTargetConfig{}
	s := strings.Split(string(data), separator)
	for i := range s {
		c := transportTargetConfig{}
		if err := yaml.Unmarshal([]byte(s[i]), &c); err != nil {
			return nil, err
		}
		b, err := yaml.Marshal(c.Spec)
		if err != nil {
			return nil, err
		}
		c.SpecBytes = b
		// check if the name is already set; require unique names
		if _, exists := interimConfigs[c.Name]; exists {
			return nil, fmt.Errorf("defined name '%s' not unique; the name is specified more than once", c.Name)
		}
		interimConfigs[c.Name] = &c
		log.Debug().Msgf("carbonaut agent %s configuration of type %s with spec length %d decoded", c.Version, c.Type, len(b))
	}
	log.Info().Msgf("target configurations detected: %d", len(interimConfigs))
	return interimConfigs, err
}

func execTargetConfig(commsChan chan CommsChannel, currentRunningProcMap map[TargetName]*targetProcess, wantedProcConfigMap map[TargetName]*transportTargetConfig) (map[TargetName]*targetProcess, error) {
	tProcessesNew := map[TargetName]*targetProcess{}
	// stop old targets configs

	// It might be the case that the wanted configured target might already be running, and therefore its not necessary to stop the process and start them again...
	// checking:
	// 1. Does the wanted process exist already in the current running process-map
	// 2. the current process is running
	// 3. if the configuration of the current process-map is the same as the wanted one
	for t := range currentRunningProcMap {
		if currentRunningProcMap[t].status == stateRunning {
			if _, exists := wantedProcConfigMap[t]; exists && reflect.DeepEqual(currentRunningProcMap[t].cfg, wantedProcConfigMap[t]) {
				tProcessesNew[t] = currentRunningProcMap[t]
				log.Debug().Msgf("target '%s' is already running", t)
			} else {
				stopTarget(currentRunningProcMap, t)
			}
		}
		delete(currentRunningProcMap, t)
	}

	// start new target configs
	for t := range wantedProcConfigMap {
		tProc, err := startTarget(commsChan, wantedProcConfigMap[t])
		if err != nil {
			return nil, err
		}
		tProcessesNew[t] = tProc
		log.Debug().Msgf("target '%s' freshly created and started", tProc.cfg.Name)
	}
	return tProcessesNew, nil
}

func startTarget(commsChan chan CommsChannel, cfg *transportTargetConfig) (*targetProcess, error) {
	interruptChan := make(chan int)
	for j := range targetImplementations {
		if targetImplementations[j].GetTargetType() != cfg.Type {
			continue
		}
		c, err := targetImplementations[j].UnmarshalSpec(cfg.SpecBytes)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("start '%s' Target '%s' with configuration: %v", cfg.Type, cfg.Name, c)

		go func(impl CarbonautAgentTarget, config any, c chan int, n TargetName, t string) {
			go func() {
				// register metrics
				registry := prometheus.NewRegistry()
				if err = impl.Register(config, registry); err != nil {
					commsChan <- CommsChannel{
						Action:  actionError,
						Name:    cfg.Name,
						Details: fmt.Sprintf("register target %s threw an error", cfg.Name),
						err:     err,
					}
				}
				// start serving registered metrics
				if err = servePrometheusMetricsEndpoint(&promServeConfig{
					ServerReadTimeout:  promServerReadTimeout,
					ServerWriteTimeout: promServerWriteTimeout,
					Port:               cfg.Port,
					Registry:           registry,
				}); err != nil {
					commsChan <- CommsChannel{
						Action:  actionError,
						Name:    cfg.Name,
						Details: fmt.Sprintf("serving target %s threw an error", cfg.Name),
						err:     err,
					}
				}
			}()
			// block; wait for call to stop target
			<-c
			log.Info().Msgf("STOPPING target %s of time %s", n, t)
		}(targetImplementations[j], c, interruptChan, cfg.Name, cfg.Type)
		break // target found; don't check other targets
	}
	return &targetProcess{
		interruptChan: interruptChan,
		status:        stateRunning,
		cfg:           *cfg,
	}, nil
}

func stopTarget(tProcesses map[TargetName]*targetProcess, name TargetName) map[TargetName]*targetProcess {
	tProcesses[name].interruptChan <- 1
	tProcesses[name] = &targetProcess{
		interruptChan: nil,
		status:        stateStopped,
		cfg:           tProcesses[name].cfg,
	}
	return tProcesses
}

func servePrometheusMetricsEndpoint(cfg *promServeConfig) error {
	log.Debug().Msgf("start prometheus on port %d", cfg.Port)
	http.Handle("/metrics", promhttp.HandlerFor(cfg.Registry, promhttp.HandlerOpts{Registry: cfg.Registry}))
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ReadTimeout:  cfg.ServerReadTimeout,
		WriteTimeout: cfg.ServerWriteTimeout,
	}
	return s.ListenAndServe()
}
