package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/handler/server"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	log := logger.FromContext(ctx)
	log.Info("Starting game server...")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	srv := server.NewServer()
	srv.RegisterHandlers(app)

	go srv.Run(ctx)

	port := "4242"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	go func() {
		defer cancel()
		if err := app.Listen(":" + port); err != nil {
			log.Errorf("Error running server: %s", err)
		}
	}()

	<-ctx.Done()
	if err := app.ShutdownWithTimeout(time.Second); err != nil {
		log.Errorf("Error shutting down server: %s", err)
	}
	log.Info("Bye!")
}
