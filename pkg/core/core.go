package core

import (
	"context"
	"fmt"

	"github.com/apelletant/upfluence-tt/pkg/domain"
)

type App struct {
	deps *domain.Dependencies
}

func NewApp(deps *domain.Dependencies) (*App, error) {
	if err := deps.Validate(); err != nil {
		return nil, fmt.Errorf("deps.validate")
	}

	app := &App{
		deps: deps,
	}
	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	return nil
}
