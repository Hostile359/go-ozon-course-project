package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host            string        `envconfig:"HOST" default:"localhost"`
	ValidGrpcPort   string        `envconfig:"VALID_GRPC_PORT" default:":8081"`
	ValidHttpPort   string        `envconfig:"VALID_HTTP_PORT" default:":8080"`
	DBGrpcPort      string        `envconfig:"DB_GRPC_PORT" default:":8082"`
	Brokers         string        `envconfig:"BROKERS" default:"localhost:9095"`
	RedisHost       string        `envconfig:"REDIS_HOST" default:"localhost:6379"`
}

func NewConfig() (*Config, error) {
	var isLocal bool

	flag.BoolVar(&isLocal, "local", false, "Use env vars from 'local.env' (should be in root)")
	flag.Parse()

	if isLocal {
		if err := godotenv.Load("local.env"); err != nil {
			return nil, err
		}
	}

	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
