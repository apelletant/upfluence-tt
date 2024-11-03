package clientmock

import (
	"time"

	"github.com/apelletant/upfluence-tt/pkg/domain"
)

var _ domain.Client = (*ClientMock)(nil)

type ClientMock struct {
	msgs []*domain.Message
}

func New(msgs []*domain.Message) *ClientMock {
	return &ClientMock{
		msgs: msgs,
	}
}

func (cm *ClientMock) Receive(_ time.Duration, msgChan chan *domain.Message) error {
	for _, msg := range cm.msgs {
		msgChan <- msg
	}
	return nil
}
