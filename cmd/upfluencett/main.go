package main

import (
	"log"

	"github.com/apelletant/upfluence-tt/pkg/core"
	"github.com/apelletant/upfluence-tt/pkg/domain"
	"github.com/apelletant/upfluence-tt/pkg/transport/http/client"
	"github.com/apelletant/upfluence-tt/pkg/transport/http/server"
	"golang.org/x/net/context"
)

func main() {
	RunApp(context.Background())
}

func RunApp(ctx context.Context) {
	cfg := parseConfig()
	if err := cfg.validate(); err != nil {
		log.Print(err.Error())

		return
	}

	scfg := &server.Config{
		Port: cfg.ServerPort,
	}

	ccfg := &client.Config{
		URL: cfg.URL,
	}

	client, err := client.New(ccfg)
	if err != nil {
		log.Print(err)

		return
	}

	server, err := server.New(scfg)
	if err != nil {
		log.Print(err)

		return
	}

	deps := &domain.Dependencies{
		Client: client,
		Server: server,
	}

	app, err := core.NewApp(deps)
	if err != nil {
		log.Print(err)

		return
	}

	server.AddApp(app)

	if err := app.Run(ctx); err != nil {
		log.Print(err)

		return
	}
}
