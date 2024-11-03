package servermock

import (
	"context"

	"github.com/apelletant/upfluence-tt/pkg/domain"
)

var _ domain.Server = (*ServerMock)(nil)

type ServerMock struct{}

func New() *ServerMock {
	return &ServerMock{}
}

func (sm *ServerMock) Run(_ context.Context) error {
	return nil
}
