package space

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

type Vec2F struct {
	X float64
	Y float64
}
