package server

import (
	"context"
	"fmt"
	"github.com/carbonaut/pkg/api"
	"github.com/carbonaut/pkg/api/model"
	"github.com/carbonaut/pkg/api/util"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"strings"
	"time"

	promApi "github.com/prometheus/client_golang/api"
	promV1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type ApiServer struct {
	api.UnimplementedEmissionDataServer
	prometheus promV1.API
}

// New api server for the gRPC API
func New(prometheusAddress string) (ApiServer, error) {
    // TODO: read this from config
	client, err := promApi.NewClient(promApi.Config{
        Address: prometheusAddress,
	})

	if err != nil {
		return ApiServer{}, err
	}

	return ApiServer{
		prometheus: promV1.NewAPI(client),
	}, nil
}

func (server *ApiServer) ListServices(ctx context.Context, in *api.ListServicesForProjectRequest) (*api.ListServicesForProjectResponse, error) {
	services := make([]string, 0)

	currentTime := time.Now().UTC()
	g, ctx := errgroup.WithContext(ctx)

	for _, providerName := range model.PROVIDER_name {
		g.Go(func() error {
			result, _, err := server.prometheus.Series(ctx, []string{strings.ToLower(providerName)}, currentTime.Add(-time.Minute*10), currentTime)
			if err != nil {
				return err
			}
			for _, labelset := range result {
				fmt.Printf("%#v", labelset)
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Err(err).Send()
		return &api.ListServicesForProjectResponse{
			Services: make([]string, 0),
			Status: &util.Status{
				Code:    util.Code_INTERNAL,
				Message: "Error while retrieving data from database",
			},
		}, err
	}

	return &api.ListServicesForProjectResponse{
		Services: services,
		Status: &util.Status{
			Code:    util.Code_OK,
			Message: "",
			Details: nil,
		},
	}, nil
}
