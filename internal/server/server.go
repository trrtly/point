package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

// A Server defines parameters for running an HTTP server.
type Server struct {
	Addr string
	App  *fiber.App
}

// ListenAndServe initializes a server to respond to HTTP network requests.
func (s Server) ListenAndServe() error {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = s.App.Shutdown()
	}()

	return s.App.Listen(s.Addr)
}
