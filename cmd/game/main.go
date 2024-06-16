package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mymmrac/hide-and-seek/pkg/game"
	_ "github.com/mymmrac/hide-and-seek/pkg/logger"
)

func main() {
	log.Info("Starting game...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	gameInstance := game.NewGame(ctx, cancel)
	if err := gameInstance.Init(); err != nil {
		log.Errorf("Error initializing game: %s", err)
		return
	}

	go runGame(gameInstance, cancel)

	<-ctx.Done()
	gameInstance.Shutdown()

	log.Info("Bye!")
}

func runGame(gameInstance *game.Game, cancel context.CancelFunc) {
	defer cancel()

	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Errorf("Error running game: %s", err)
	}
}
