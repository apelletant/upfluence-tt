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
	if err := a.deps.Server.Run(ctx); err != nil {
		return fmt.Errorf("a.deps.Server.Run: %w", err)
	}

	return nil
}

func (a *App) RunQuery(dimension string, duration string) (map[string]int, error) {
	log.Printf("Analyzing %s for %s\n", dimension, duration)

	ttl, err := time.ParseDuration(duration)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}

	msgChan := make(chan *domain.Message)

	go func() {
		defer close(msgChan)

		if err := a.deps.Client.Receive(ttl, msgChan); err != nil {
			log.Print(err)

			return
		}
	}()

	res := make(map[string]int)

	for msg := range msgChan {
		if msg.Err != nil {
			log.Print(err)

			continue
		}

		if res["total_posts"] == 0 {
			res["minimum_timestamp"] = msg.Data.Timestamp
			res["maximum_timestamp"] = msg.Data.Timestamp
		}

		if msg.Data.Timestamp > res["maximum_timestamp"] {
			res["maximum_timestamp"] = msg.Data.Timestamp
		}

		if msg.Data.Timestamp < res["minimum_timestamp"] {
			res["minimum_timestamp"] = msg.Data.Timestamp
		}

		switch dimension {
		case domain.Likes:
			if msg.Data.Likes != nil {
				res["avg_likes"] += *msg.Data.Likes
				res["total_posts"]++
			}
		case domain.Retweets:
			if msg.Data.Retweets != nil {
				res["avg_retweets"] += *msg.Data.Retweets
				res["total_posts"]++
			}
		case domain.Comments:
			if msg.Data.Comments != nil {
				res["avg_comments"] += *msg.Data.Comments
				res["total_posts"]++
			}
		case domain.Favorites:
			if msg.Data.Favorites != nil {
				res["avg_favorites"] += *msg.Data.Favorites
				res["total_posts"]++
			}
		default:
			return nil, domain.ErrDimensionUnknown
		}
	}

	res["avg_"+dimension] = res["avg_"+dimension] / res["total_posts"]

	return res, nil
}
