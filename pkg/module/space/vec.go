package space

import (
	"math"
	"math/rand/v2"
)

type Vec2I struct {
	X int
	Y int
}

func (v Vec2I) ToF() Vec2F {
	return Vec2F{
		X: float64(v.X),
		Y: float64(v.Y),
	}
}

func (v Vec2I) Add(u Vec2I) Vec2I {
	return Vec2I{
		X: v.X + u.X,
		Y: v.Y + u.Y,
	}
}

func (v Vec2I) ScaleVec(u Vec2I) Vec2I {
	return Vec2I{
		X: v.X * u.X,
		Y: v.Y * u.Y,
	}
}

type Vec2F struct {
	X float64
	Y float64
}

func RandomVec2F() Vec2F {
	return Vec2F{
		X: rand.Float64(),
		Y: rand.Float64(),
	}
}

func (v Vec2F) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec2F) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2F) Len() float64 {
	return math.Sqrt(v.Len2())
}

func (v Vec2F) Norm() Vec2F {
	l := v.Len()
	return Vec2F{
		X: v.X / l,
		Y: v.Y / l,
	}
}

func (v Vec2F) Add(u Vec2F) Vec2F {
	return Vec2F{
		X: v.X + u.X,
		Y: v.Y + u.Y,
	}
}

func (v Vec2F) Dist2(u Vec2F) float64 {
	return (v.X-u.X)*(v.X-u.X) + (v.Y-u.Y)*(v.Y-u.Y)
}

func (v Vec2F) AddX(x float64) Vec2F {
	return Vec2F{
		X: v.X + x,
		Y: v.Y,
	}
}

func (v Vec2F) AddY(y float64) Vec2F {
	return Vec2F{
		X: v.X,
		Y: v.Y + y,
	}
}

func (v Vec2F) Sub(u Vec2F) Vec2F {
	return Vec2F{
		X: v.X - u.X,
		Y: v.Y - u.Y,
	}
}

func (v Vec2F) Scale(n float64) Vec2F {
	return Vec2F{
		X: v.X * n,
		Y: v.Y * n,
	}
}

func (v Vec2F) OX() Vec2F {
	return Vec2F{
		X: v.X,
		Y: 0,
	}
}

func (v Vec2F) OY() Vec2F {
	return Vec2F{
		X: 0,
		Y: v.Y,
	}
}
