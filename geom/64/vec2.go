package geom

import (
	"math"
	"math/rand"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Ori2() Ori2 {
	return Ori2{v.X, v.Y, 0}
}

func (a Vec2) Plus(b Vec2) Vec2 {
	return Vec2{a.X + b.X, a.Y + b.Y}
}

func (a Vec2) Minus(b Vec2) Vec2 {
	return Vec2{a.X - b.X, a.Y - b.Y}
}

func (a Vec2) Cross(b Vec2) float64 {
	return a.X*b.Y - a.Y*b.X
}

func (v Vec2) ScaledBy(f float64) Vec2 {
	return Vec2{v.X * f, v.Y * f}
}

func (v Vec2) RotatedBy(radians float64) Vec2 {
	sin := math.Sin(radians)
	cos := math.Cos(radians)
	return Vec2{cos*v.X - sin*v.Y, sin*v.X + cos*v.Y}
}

func (v Vec2) Len2() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Len() float64 {
	return math.Sqrt(v.Len2())
}

func (v Vec2) Normal() Vec2 {
	len := v.Len()
	if len == 0 || math.IsNaN(v.X) || math.IsNaN(v.Y) {
		return Vec2{}
	}

	switch {
	case math.IsInf(v.X, 0) && math.IsInf(v.Y, 0):
		return Vec2{}
	case math.IsInf(v.X, 1):
		return Vec2{1, 0}
	case math.IsInf(v.X, -1):
		return Vec2{-1, 0}
	case math.IsInf(v.Y, 1):
		return Vec2{0, 1}
	case math.IsInf(v.Y, -1):
		return Vec2{0, -1}
	}

	return v.ScaledBy(1 / len)
}

func (a *Vec2) PlusEquals(b Vec2) {
	a.X += b.X
	a.Y += b.Y
}

func Vec2Rand(r Rect) Vec2 {
	return Vec2{
		rand.Float64()*r.Width() + r.Min.X,
		rand.Float64()*r.Height() + r.Min.Y,
	}
}

func Vec2RandNormal() Vec2 {
	theta := rand.Float64() * 2 * math.Pi
	return Vec2{
		math.Sin(theta),
		math.Cos(theta),
	}
}
