package game

import (
	"fmt"
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkgray)

	for _, lvl := range g.world.Levels {
		for _, tile := range lvl.Tiles {
			tileset := g.defs.Tilesets[tile.TilesetID]
			tileDef := tileset.Tiles[tile.TileID]
			tileImage := g.tilesets[tile.TilesetID].SubImage(image.Rect(
				tileDef.X, tileDef.Y,
				tileDef.X+tileset.TileSize.X, tileDef.Y+tileset.TileSize.Y,
			))

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(lvl.Pos.X), float64(lvl.Pos.Y))
			op.GeoM.Translate(float64(tile.Pos.X), float64(tile.Pos.Y))
			screen.DrawImage(tileImage.(*ebiten.Image), op)
		}
	}

	const pad = 4
	vector.DrawFilledRect(
		screen,
		float32(g.world.Spawn.X)+pad,
		float32(g.world.Spawn.Y)+pad,
		32-pad*2,
		32-pad*2,
		colornames.Orange,
		true,
	)

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
