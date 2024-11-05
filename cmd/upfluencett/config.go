package main

import (
	"errors"
	"flag"
)

var (
	ErrURLMissing  = errors.New("upfluence API url not set")
	ErrPortMissing = errors.New("server port not set")
)

type Config struct {
	URL        string
	ServerPort int
}

func (cfg Config) validate() error {
	if cfg.URL == "" {
		return ErrURLMissing
	}

	if cfg.ServerPort == -1 {
		return ErrPortMissing
	}

	return nil
}

func parseConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.URL, "upfluence-url", "localhost:80", "url use to retrieve data")
	flag.IntVar(&cfg.ServerPort, "server-port", -1, "port use main server")

	flag.Parse()

	return cfg
}
