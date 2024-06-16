package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mymmrac/hide-and-seek/pkg/handler/game"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	log := logger.FromContext(ctx)
	log.Info("Starting game...")

	gameInstance := game.NewGame(ctx, cancel)
	if err := gameInstance.Init(); err != nil {
		log.Errorf("Error initializing game: %s", err)
		return
	}

	go func() {
		defer cancel()
		if err := ebiten.RunGame(gameInstance); err != nil {
			log.Errorf("Error running game: %s", err)
		}
	}()

	<-ctx.Done()
	gameInstance.Shutdown()

	log.Info("Bye!")
}
