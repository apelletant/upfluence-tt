package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/apelletant/upfluence-tt/pkg/domain"
)

var _ domain.Client = (*Client)(nil)

var (
	ErrURLNotSet      = errors.New("url cannot be emoty")
	ErrRetreivingData = errors.New("while parsing input data")
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

func (c *Client) Receive(ttl time.Duration, msgChan chan *domain.Message) error {
	req, err := http.NewRequest("GET", c.cfg.URL, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("c.Do: %w", err)
	}

	defer res.Body.Close()

	timer := time.NewTimer(ttl)
	defer timer.Stop()

	scanner := bufio.NewScanner(res.Body)

	done := make(chan struct{})

	go func() {
		<-timer.C
		close(done)
	}()

	for {
		select {
		case <-done:
			return nil
		default:
			if scanner.Scan() {
				line := scanner.Text()
				if len(line) > 6 && line[:6] == "data: " {
					msg := ExtractMessage(line[6:]) // removinfgg "data: "
					msgChan <- msg
				}
			} else if err := scanner.Err(); err != nil {
				log.Print("Scanner error", err)
			}
		}
	}
}

func ExtractMessage(rawMessage string) *domain.Message {
	msg, err := parseData(rawMessage)
	if err != nil {
		return &domain.Message{
			Data: nil,
			Err:  err,
		}
	}

	return &domain.Message{
		Data: msg,
		Err:  nil,
	}
}

func parseData(data string) (*domain.MsgData, error) {
	var obj map[string]interface{}

	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	for _, v := range obj {
		if valueMap, ok := v.(map[string]interface{}); ok {
			msgData := &domain.MsgData{
				Likes:     ToIntPointer(valueMap[domain.Likes]),
				Comments:  ToIntPointer(valueMap[domain.Comments]),
				Favorites: ToIntPointer(valueMap[domain.Favorites]),
				Retweets:  ToIntPointer(valueMap[domain.Retweets]),
			}

			if ts, ok := valueMap["timestamp"].(float64); ok {
				msgData.Timestamp = int(ts)
			}

			// should only have one item
			return msgData, nil
		}
	}

	return nil, ErrRetreivingData
}

func ToIntPointer(itp interface{}) *int {
	if v, ok := itp.(float64); ok {
		res := int(v)

		return &res
	}

	if v, ok := itp.(int); ok {
		return &v
	}

	return nil
}
