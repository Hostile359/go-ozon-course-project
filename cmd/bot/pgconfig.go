package main

import (
	"flag"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host            string        `envconfig:"HOST" default:"localhost"`
	Port            int           `envconfig:"PORT" default:"5432"`
	User            string        `envconfig:"DB_USER" required:"true"`
	Password        string        `envconfig:"DB_PASSWORD" required:"true"`
	DBname          string        `envconfig:"DBNAME" required:"true"`
	MaxConnIdleTime time.Duration `envconfig:"MAX_CONN_IDLE_TIME" default:"1m" required:"true"`
	MaxConnLifetime time.Duration `envconfig:"MAX_CONN_LIFE_TIME" default:"1h" required:"true"`
	MinConns        int32         `envconfig:"MIN_CONNS" default:"2" required:"true"`
	MaxConns        int32         `envconfig:"MAX_CONNS" default:"4" required:"true"`
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
