package game

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkgray)

	vector.DrawFilledRect(
		screen,
		float32(g.player.Pos.X),
		float32(g.player.Pos.Y),
		32,
		32,
		colornames.Blue,
		true,
	)
	ebitenutil.DebugPrintAt(screen, g.player.Name, int(g.player.Pos.X), int(g.player.Pos.Y))

	g.players.ForEach(func(_ uint64, player *Player) bool {
		vector.DrawFilledRect(
			screen,
			float32(player.Pos.X),
			float32(player.Pos.Y),
			32,
			32,
			colornames.Green,
			true,
		)
		ebitenutil.DebugPrintAt(screen, player.Name, int(player.Pos.X), int(player.Pos.Y))
		return true
	})

	debugDraw(screen,
		fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()),
		fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()),
		"",
		fmt.Sprintf("Connected: %t", g.connected),
	)
}

func debugDraw(screen *ebiten.Image, lines ...string) {
	ebitenutil.DebugPrint(screen, strings.Join(lines, "\n"))
}
