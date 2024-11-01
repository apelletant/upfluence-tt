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
	Receive(ttr time.Duration, msgChan chan *Message) error
}

type Message struct {
	Data map[string]interface{}
	Err  error
}

type Response struct {
	TotalPosts int `json:"total_posts"`
	MinTS      int `json:"minimum_timestamp"`
	MxwTS      int `json:"maximum_timestamp"`
	AvgLikes   int `json:"avg_likes"`
}
