package geom

import (
	"math"
	"math/rand"
)

type Vec2 struct {
	X, Y float32
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

func (a Vec2) Cross(b Vec2) float32 {
	return a.X*b.Y - a.Y*b.X
}

func (v Vec2) ScaledBy(f float32) Vec2 {
	return Vec2{v.X * f, v.Y * f}
}

func (v Vec2) RotatedBy(radians float32) Vec2 {
	radians64 := float64(radians)
	sin := float32(math.Sin(radians64))
	cos := float32(math.Cos(radians64))

	return Vec2{cos*v.X - sin*v.Y, sin*v.X + cos*v.Y}
}

func (v Vec2) Len2() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Len() float32 {
	return float32(math.Sqrt(float64(v.Len2())))
}

func (v Vec2) Normal() Vec2 {
	x64 := float64(v.X)
	y64 := float64(v.Y)

	len := v.Len()
	if len == 0 || math.IsNaN(x64) || math.IsNaN(y64) {
		return Vec2{}
	}

	switch {
	case math.IsInf(x64, 0) && math.IsInf(y64, 0):
		return Vec2{}
	case math.IsInf(x64, 1):
		return Vec2{1, 0}
	case math.IsInf(x64, -1):
		return Vec2{-1, 0}
	case math.IsInf(y64, 1):
		return Vec2{0, 1}
	case math.IsInf(y64, -1):
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
		rand.Float32()*r.Width() + r.Min.X,
		rand.Float32()*r.Height() + r.Min.Y,
	}
}

func Vec2RandNormal() Vec2 {
	theta := rand.Float64() * 2 * math.Pi
	return Vec2{
		float32(math.Sin(theta)),
		float32(math.Cos(theta)),
	}
}
