package main

import (
	"point/cmd/server/config"
	"point/internal/handler/api"
	"point/internal/handler/health"
	"point/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

type (
	healthzHandler *fiber.App
)

// wire set for loading the server.
var serverSet = wire.NewSet(
	api.New,
	provideRouter,
	provideServer,
)

// provideRouter is a Wire provider function that returns a
// router that is serves the provided handlers.
func provideRouter(api api.Server) *fiber.App {
	r := fiber.New()
	r.Mount("/healthz", health.Handler())
	r.Mount("/api", api.Handler())
	return r
}

// provideServer is a Wire provider function that returns an
// http server that is configured from the environment.
func provideServer(app *fiber.App, config config.Config) *server.Server {
	return &server.Server{
		Addr: config.Server.Port,
		Host: config.Server.Host,
		App:  app,
	}
}