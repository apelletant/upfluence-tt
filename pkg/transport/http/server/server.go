package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/apelletant/upfluence-tt/pkg/core"
	"github.com/apelletant/upfluence-tt/pkg/domain"
	echo "github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

var _ domain.Server = (*Server)(nil)

var (
	ErrPortNotSet = errors.New("port should be set")
	ErrServerNil  = errors.New("server cannot be nil")
	ErrAppNil     = errors.New("app cannot be nil")
)

type Server struct {
	cfg *Config
	e   *echo.Echo
	app *core.App
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
		cfg: cfg,
		e:   e,
	}

	s.e.GET("/", s.handleDefault)
	s.e.GET("/analysis", s.analysis)

	return s, nil
}

func (s *Server) AddApp(app *core.App) {
	s.app = app
}

func (s *Server) Run(ctx context.Context) error {
	errG, errCtx := errgroup.WithContext(ctx)

	errG.Go(func() error {
		err := s.e.Start(fmt.Sprintf(":%d", s.cfg.Port))
		if err != nil {
			return fmt.Errorf("s.e.Run: %w", err)
		}

		return nil
	})

	<-errCtx.Done()

	if err := s.e.Close(); err != nil {
		return fmt.Errorf("s.e.close: %w", err)
	}

	if err := errG.Wait(); err != nil {
		return fmt.Errorf("errG.Wait: %w", err)
	}

	return nil
}

func (s *Server) Close() error {
	if err := s.e.Close(); err != nil {
		return fmt.Errorf("s.e.Close")
	}

	return nil
}

func (s *Server) handleDefault(ectx echo.Context) error {
	return ectx.JSON(s.buildResponse(http.StatusNotFound)) //nolint:wrapcheck
}

func (s *Server) analysis(ectx echo.Context) error {
	duration := ectx.QueryParam("duration")
	dimension := ectx.QueryParam("dimension")

	if duration == "" {
		return ectx.JSON(s.buildResponseWithMessage(http.StatusBadRequest, "missing duration queryparam")) //nolint:wrapcheck
	}

	if dimension == "" {
		return ectx.JSON(s.buildResponseWithMessage(http.StatusBadRequest, "missing dimension queryparam")) //nolint:wrapcheck
	}

	result, err := s.app.RunQuery(dimension, duration)
	if err != nil {
		if errors.Is(err, domain.ErrDimensionUnknown) {
			return ectx.JSON(s.buildResponseWithMessage(http.StatusBadRequest, err.Error())) //nolint:wrapcheck
		}

		return ectx.JSON(s.buildResponseWithMessage(http.StatusInternalServerError, err.Error())) //nolint:wrapcheck
	}

	b, err := json.Marshal(result)
	if err != nil {
		return ectx.JSON(s.buildResponseWithMessage(http.StatusInternalServerError, err.Error())) //nolint:wrapcheck
	}

	return ectx.JSON(http.StatusOK, string(b)) //nolint:wrapcheck
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

func (s *Server) buildResponseWithMessage(statusCode int, message string) (int, *Response) {
	return statusCode, &Response{
		StatusCode: statusCode,
		Message:    message,
	}
}
