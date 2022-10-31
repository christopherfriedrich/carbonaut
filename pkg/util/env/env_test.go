/*
Copyright (c) 2022 CARBONAUT AUTHOR

Licensed under the MIT license: https://opensource.org/licenses/MIT
Permission is granted to use, copy, modify, and redistribute the work.
Full license information available in the project LICENSE file.
*/

package env_test

import (
	"os"
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/carbonaut/pkg/util/env"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEnv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Env Suite")
}

var _ = Describe("Env", func() {
	tmpEnvK1 := "CARBONAUT_TEMP"
	tmpEnvV1 := "ABC"

	tmpEnvK2 := "CARBONAUT_TEMP_2"
	tmpEnvV2 := "ABC_2"
	BeforeEach(func() {
		if err := os.Setenv(tmpEnvK1, tmpEnvV1); err != nil {
			log.Fatal().Err(err)
		}
	})
	AfterSuite(func() {
		// unset the environment variable
		if err := os.Setenv(tmpEnvK2, ""); err != nil {
			log.Fatal().Err(err)
		}
		if err := os.Setenv(tmpEnvK1, ""); err != nil {
			log.Fatal().Err(err)
		}
	})

	Context("an environment variable is set", func() {
		It("should detect an existing environment variable", func() {
			Expect(env.IsSet(tmpEnvK1)).To(BeTrue())
		})
	})
	Context("an environment variable is not set", func() {
		It("should detect that the environment variable is not set", func() {
			Expect(env.IsSet(tmpEnvK2)).To(BeFalse())
		})
	})

	Context("an environment variable which gets set at runtime", func() {
		It("should return the value of the environment variable", func() {
			Expect(env.Default(tmpEnvK2, tmpEnvV2)).To(Equal(tmpEnvV2))
		})
		It("should detect that the environment variable is set", func() {
			Expect(env.IsSet(tmpEnvK2)).To(BeTrue())
		})
	})
})
