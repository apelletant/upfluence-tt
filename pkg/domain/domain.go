package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrServerNil = errors.New("server cannot be nil")
	ErrClientNil = errors.New("client cannot be nil")
)

type Dependencies struct {
	Server Server
	Client Client
}

func (d *Dependencies) Validate() error {
	if d.Server == nil {
		return ErrServerNil
	}

	if d.Client == nil {
		return ErrClientNil
	}

	return nil
}

type Server interface {
	Run(ctx context.Context) error
}

type Client interface {
	Receive(ttr time.Duration) error
}
