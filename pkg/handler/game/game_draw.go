package game

import (
	"cmp"
	"fmt"
	"image"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/exp/slices"
	"golang.org/x/image/colornames"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkslategray)
	g.worldImg.Fill(colornames.Darkslategray)

	for _, lvl := range g.world.Levels {
		for _, tile := range lvl.Tiles {
			tileset := g.defs.Tilesets[tile.TilesetID]
			tileDef := tileset.Tiles[tile.TileID]
			tileImage := g.tilesets[tile.TilesetID].SubImage(image.Rect(
				tileDef.Pos.X, tileDef.Pos.Y,
				tileDef.Pos.X+tileDef.Size.X, tileDef.Pos.Y+tileDef.Size.Y,
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

	// Draw spawn
	// const pad = 4
	// vector.DrawFilledRect(
	// 	g.worldImg,
	// 	float32(g.world.Spawn.X)+pad,
	// 	float32(g.world.Spawn.Y)+pad,
	// 	32-pad*2,
	// 	32-pad*2,
	// 	colornames.Orange,
	// 	true,
	// )

	var drawCalls []DrawCall

	for _, lvl := range g.world.Levels {
		for _, entity := range lvl.Entities {
			def := g.defs.Entities[entity.EntityID]

			tileset := g.defs.Tilesets[def.TilesetID]
			tileDef := tileset.Tiles[def.TileID]
			tileImage := g.tilesets[def.TilesetID].SubImage(image.Rect(
				tileDef.Pos.X, tileDef.Pos.Y,
				tileDef.Pos.X+tileDef.Size.X, tileDef.Pos.Y+tileDef.Size.Y,
			))

			op := &ebiten.DrawImageOptions{}
			pos := entity.Pos.ToF()
			op.GeoM.Translate(pos.X, pos.Y)
			drawCalls = append(drawCalls, DrawCall{
				Y: pos.Y,
				F: func() {
					g.worldImg.DrawImage(tileImage.(*ebiten.Image), op)
				},
			})
		}
	}

	drawCalls = append(drawCalls, g.player.Draw(g.worldImg, g.playerSpriteSheet))

	g.players.ForEach(func(_ uint64, player *Player) bool {
		drawCalls = append(drawCalls, player.Draw(g.worldImg, g.playerSpriteSheet))
		return true
	})

	slices.SortStableFunc(drawCalls, func(a, b DrawCall) int {
		return cmp.Compare(a.Y, b.Y)
	})
	for _, drawCall := range drawCalls {
		drawCall.F()
	}

	// Draw colliders
	// for _, obj := range g.cw.Objects() {
	// 	vector.StrokeRect(
	// 		g.worldImg,
	// 		float32(obj.Position().X),
	// 		float32(obj.Position().Y),
	// 		float32(obj.Size().X),
	// 		float32(obj.Size().Y),
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

type DrawCall struct {
	Y float64
	F func()
}
