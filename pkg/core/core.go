package core

import (
	"context"
	"fmt"
	"log"
	"time"

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
	a.deps.Server.Run(ctx)
	return nil
}

func (a *App) RunQuery(dimension string, duration string) error {
	ttl, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("time.Parse: %w", err)
	}

	msgChan := make(chan *domain.Message)
	defer close(msgChan)

	go func() {
		a.deps.Client.Receive(ttl, msgChan)
	}()

	res := &domain.Response{}

	for msg := range msgChan {
		if msg.Err != nil {
			log.Print(err)

			continue
		}

		/*
		for _, v := range msg.Data {
			//fmt.Println("key", k, "value", v)
			switch v.(type) {
			case domain.Instagram:
				fmt.Println("instagram")
			case domain.Tiktok:
				fmt.Println("tiktok")
			case domain.Twitch:
				fmt.Println("twitch")
			default:
				fmt.Println("default")
			}
*/
			res.TotalPosts++
		}
	}

	return nil
}
