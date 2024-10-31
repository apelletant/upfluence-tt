package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/apelletant/upfluence-tt/pkg/domain"
	echo "github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

var _ domain.Server = (*Server)(nil)

var (
	ErrPortNotSet = errors.New("port should be set")
	ErrServerNil  = errors.New("server cannot be nil")
)

type Server struct {
	port int
	e    *echo.Echo
}

type Config struct {
	Port int
}

func (cfg *Config) validate() error {
	if cfg.Port == -1 {
		return ErrPortNotSet
	}

	return nil
}

func New(cfg *Config) (*Server, error) {
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("cfg.validate: %w", err)
	}

	e := echo.New()

	s := &Server{
		port: cfg.Port,
		e:    e,
	}

	s.e.GET("/", s.handleDefault)
	s.e.GET("/analysis", s.analysis)

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	errG, errCtx := errgroup.WithContext(ctx)

	errG.Go(func() error {
		err := s.e.Start(fmt.Sprintf(":%d", s.port))
		if err != nil {
			return fmt.Errorf("s.e.Run: %w", err)
		}

		return nil
	})

	<-errCtx.Done()

	if err := s.e.Close(); err != nil {
		// TODO HAndle log
		//clog.Logger.Error("s.e.Close", clog.Error(err))
	}

	if err := errG.Wait(); err != nil {
		return fmt.Errorf("errG.Wait: %w", err)
	}

	return nil
}

func (s *Server) handleDefault(ectx echo.Context) error {
	return ectx.JSON(s.buildResponse(http.StatusNotFound))
}

func (s *Server) analysis(ectx echo.Context) error {
	return ectx.JSON(s.buildResponse(http.StatusOK))
}

type Response struct {
	Message    string
	StatusCode int
}

func (s *Server) buildResponse(statusCode int) (int, *Response) {
	return statusCode, &Response{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
	}
}
