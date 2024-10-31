package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/apelletant/upfluence-tt/pkg/domain"
)

var _ domain.Client = (*Client)(nil)

var (
	ErrURLNotSet = errors.New("url cannot be emoty")
)

type Client struct {
	cfg *Config
}

type Config struct {
	URL string
}

func (cfg *Config) validate() error {
	if cfg.URL == "" {
		return ErrURLNotSet
	}

	return nil
}

func New(cfg *Config) (*Client, error) {
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("cfg.validate: %w", err)
	}

	return &Client{
		cfg: cfg,
	}, nil
}

func (s *Client) Receive(ttr time.Duration) error {
	return nil
}
