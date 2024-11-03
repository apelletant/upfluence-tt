package domain

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrServerNil        = errors.New("server cannot be nil")
	ErrClientNil        = errors.New("client cannot be nil")
	ErrDimensionUnknown = errors.New("unknown dimension")
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

// using pointer to get nil instead of default value 0
// which will be used remove message for result computing
// if it does not contains the dimension needed.
type MsgData struct {
	Likes     *int `json:"likes,omitempty"`
	Comments  *int `json:"comments,omitempty"`
	Favorites *int `json:"favorites,omitempty"`
	Retweets  *int `json:"retweets,omitempty"`
	Timestamp int  `json:"timestamp"`
}

func (msg MsgData) String() string {
	s := ""

	if msg.Favorites != nil {
		fmt.Sprintf("%s favorites: %d", s)
	}
	if msg.Likes != nil {
		fmt.Sprintf("%s likes: %d", s)
	}
	if msg.Comments != nil {
		fmt.Sprintf("%s comments: %d", s)
	}
	if msg.Retweets != nil {
		fmt.Sprintf("%s retweets: %d", s)
	}

	return fmt.Sprintf("%s timestamp: %d", s, msg.Timestamp)
}

type Message struct {
	Data *MsgData
	Err  error
}

type Response struct {
	TotalPosts   int `json:"total_posts"`
	MinTS        int `json:"minimum_timestamp"`
	MaxTS        int `json:"maximum_timestamp"`
	AvgLikes     int `json:"avg_likes,omitempty"`
	AvgComments  int `json:"avg_comments,omitempty"`
	AvgFavorites int `json:"avg_favorites,omitempty"`
	AvgRetweets  int `json:"avg_retweets,omitempty"`
}
