package cpdrawer

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
)

var (
	whiteImage    = ebiten.NewImage(3, 3)
	whiteSubImage = whiteImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	whiteImage.Fill(color.White)
}

type drawer struct {
	screen *ebiten.Image
}

func NewDrawer(screen *ebiten.Image) cp.Drawer {
	return &drawer{
		screen: screen,
	}
}

func (d drawer) DrawCircle(pos cp.Vector, _, radius float64, _, fill cp.FColor, _ any) {
	vector.DrawFilledCircle(d.screen, float32(pos.X), float32(pos.Y), float32(radius), clrFrom(fill), true)
}

func (d drawer) DrawSegment(a, b cp.Vector, fill cp.FColor, _ any) {
	vector.StrokeLine(d.screen, float32(a.X), float32(a.Y), float32(b.X), float32(b.Y), 1, clrFrom(fill), true)
}

func (d drawer) DrawFatSegment(a, b cp.Vector, radius float64, _, fill cp.FColor, _ any) {
	vector.StrokeLine(d.screen, float32(a.X), float32(a.Y), float32(b.X), float32(b.Y), float32(radius*2), clrFrom(fill), true)
}

func (d drawer) DrawPolygon(_ int, verts []cp.Vector, _ float64, _, fill cp.FColor, _ any) {
	path := vector.Path{}
	path.MoveTo(float32(verts[0].X), float32(verts[0].Y))
	for _, vs := range verts {
		path.LineTo(float32(vs.X), float32(vs.Y))
	}
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

	whiteImage.Fill(clrFrom(fill))
	d.screen.DrawTriangles(vs, is, whiteSubImage, &ebiten.DrawTrianglesOptions{
		AntiAlias: true,
	})
}

func (d drawer) DrawDot(size float64, pos cp.Vector, fill cp.FColor, _ any) {
	vector.DrawFilledCircle(d.screen, float32(pos.X), float32(pos.Y), float32(size/2), clrFrom(fill), true)
}

func (d drawer) Flags() uint {
	return 0
}

func (d drawer) OutlineColor() cp.FColor {
	return clrTo(colornames.Blue)
}

func (d drawer) ShapeColor(_ *cp.Shape, _ any) cp.FColor {
	return clrTo(colornames.Lightblue)
}

func (d drawer) ConstraintColor() cp.FColor {
	return clrTo(colornames.Orange)
}

func (d drawer) CollisionPointColor() cp.FColor {
	return clrTo(colornames.Red)
}

func (d drawer) Data() any {
	return nil
}

func clrTo(clr color.Color) cp.FColor {
	r, g, b, a := clr.RGBA()
	return cp.FColor{
		R: float32(r) / 0xff,
		G: float32(g) / 0xff,
		B: float32(b) / 0xff,
		A: float32(a) / 0xff,
	}
}

func clrFrom(clr cp.FColor) color.Color {
	return color.RGBA{
		R: uint8(clr.R * 0xff),
		G: uint8(clr.G * 0xff),
		B: uint8(clr.B * 0xff),
		A: uint8(clr.A * 0xff),
	}
}
