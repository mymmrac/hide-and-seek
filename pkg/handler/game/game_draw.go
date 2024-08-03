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
	g.worldImg.Fill(colornames.Darkgray)

	for _, lvl := range g.world.Levels {
		for _, tile := range lvl.Tiles {
			tileset := g.defs.Tilesets[tile.TilesetID]
			tileDef := tileset.Tiles[tile.TileID]
			tileImage := g.tilesets[tile.TilesetID].SubImage(image.Rect(
				tileDef.X, tileDef.Y,
				tileDef.X+tileset.TileSize.X, tileDef.Y+tileset.TileSize.Y,
			))

			op := &ebiten.DrawImageOptions{}
			pos := lvl.Pos.Add(tile.Pos).ToF()
			op.GeoM.Translate(pos.X, pos.Y)
			g.worldImg.DrawImage(tileImage.(*ebiten.Image), op)
		}
	}

	// Draw tile IDs
	// for _, lvl := range g.world.Levels {
	// 	for _, tile := range lvl.Tiles {
	// 		pos := lvl.Pos.Add(tile.Pos).Add(space.Vec2I{X: 8, Y: 8})
	// 		ebitenutil.DebugPrintAt(g.worldImg, fmt.Sprintf("%d", tile.TileID), pos.X, pos.Y)
	// 	}
	// }

	const pad = 4
	vector.DrawFilledRect(
		g.worldImg,
		float32(g.world.Spawn.X)+pad,
		float32(g.world.Spawn.Y)+pad,
		32-pad*2,
		32-pad*2,
		colornames.Orange,
		true,
	)

	vector.DrawFilledRect(
		g.worldImg,
		float32(g.player.Pos.X),
		float32(g.player.Pos.Y),
		32,
		32,
		colornames.Blue,
		true,
	)
	ebitenutil.DebugPrintAt(g.worldImg, g.player.Name, int(g.player.Pos.X), int(g.player.Pos.Y))

	g.players.ForEach(func(_ uint64, player *Player) bool {
		vector.DrawFilledRect(
			g.worldImg,
			float32(player.Pos.X),
			float32(player.Pos.Y),
			32,
			32,
			colornames.Green,
			true,
		)
		ebitenutil.DebugPrintAt(g.worldImg, player.Name, int(player.Pos.X), int(player.Pos.Y))
		return true
	})

	// Draw colliders
	// for _, obj := range g.space.Objects() {
	// 	vector.StrokeRect(
	// 		g.worldImg,
	// 		float32(obj.Position.X),
	// 		float32(obj.Position.Y),
	// 		float32(obj.Size.X),
	// 		float32(obj.Size.Y),
	// 		2,
	// 		colornames.Lightgreen,
	// 		true,
	// 	)
	// }

	g.camera.Render(g.worldImg, screen)

	// Draw world position
	// curX, curY := ebiten.CursorPosition()
	// wX, wY := g.camera.ScreenToWorld(float64(curX), float64(curY))
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.2f, %.2f", wX, wY), curX+16, curY+16)

	debugDraw(screen,
		fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()),
		fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()),
		"",
		g.camera.String(),
		"",
		fmt.Sprintf("Connected: %t", g.connected),
	)
}

func debugDraw(screen *ebiten.Image, lines ...string) {
	ebitenutil.DebugPrint(screen, strings.Join(lines, "\n"))
}
