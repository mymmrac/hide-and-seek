package camera

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mymmrac/hide-and-seek/pkg/module/space"
)

type Camera struct {
	Viewport space.Vec2F
	Position space.Vec2F
	Zoom     int
	Rotation int
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(x, y float64) (float64, float64) {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(x, y)
	} else {
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) String() string {
	return fmt.Sprintf(
		"T: %.1f, R: %d, S: %d",
		c.Position, c.Rotation, c.Zoom,
	)
}

func (c *Camera) ViewportCenter() space.Vec2F {
	return space.Vec2F{
		X: c.Viewport.X * 0.5,
		Y: c.Viewport.Y * 0.5,
	}
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position.X, -c.Position.Y)
	// We want to scale and rotate around the center of the image / screen
	m.Translate(-c.ViewportCenter().X, -c.ViewportCenter().Y)
	m.Scale(
		math.Pow(1.01, float64(c.Zoom)),
		math.Pow(1.01, float64(c.Zoom)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.ViewportCenter().X, c.ViewportCenter().Y)
	return m
}
