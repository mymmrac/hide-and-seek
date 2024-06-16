package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	_ "github.com/mymmrac/hide-and-seek/pkg/logger"
	"github.com/mymmrac/hide-and-seek/pkg/server"
)

func main() {
	log.Info("Starting game server...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	srv := server.NewServer()

	app.Get("/", websocket.New(srv.Handler))

	go runServer(app, cancel)

	<-ctx.Done()
	if err := app.ShutdownWithTimeout(time.Second); err != nil {
		log.Errorf("Error shutting down server: %s", err)
	}
	log.Info("Bye!")
}

func runServer(app *fiber.App, cancel context.CancelFunc) {
	defer cancel()

	if err := app.Listen(":4242"); err != nil {
		log.Errorf("Error running server: %s", err)
	}
}
