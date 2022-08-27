// +build integration

package config

import (
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "QA"

type Config struct {
	Host       string        `envconfig:"HOST" default:"localhost"`
	GrpcPort   string        `envconfig:"GRPC_PORT" default:":8081"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process(envPrefix, config); err != nil {
		return nil, err
	}

	return config, nil
}
