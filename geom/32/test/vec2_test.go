package geomTest

import (
	. "github.com/tadeuszjt/geom/32"
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func vec2Identical(a, b Vec2) bool {
	return floatIdentical(a.X, b.X) && floatIdentical(a.Y, b.Y)
}

func TestVec2Ori2(t *testing.T) {
	cases := []struct {
		v Vec2
		o Ori2
	}{
		{Vec2{}, Ori2{}},
		{Vec2{1, 2}, Ori2{1, 2, 0}},
		{Vec2{nInf, nan}, Ori2{nInf, nan, 0}},
	}
	for _, c := range cases {
		expected := c.o
		actual := c.v.Ori2()
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestVec2Plus(t *testing.T) {
	cases := []struct {
		A, B, result Vec2
	}{
		{Vec2{0, 0}, Vec2{0, 0}, Vec2{0, 0}},
		{Vec2{1, 2}, Vec2{0, 0}, Vec2{1, 2}},
		{Vec2{-1, 2}, Vec2{3, 4}, Vec2{2, 6}},
		{Vec2{-1, -2}, Vec2{-3, -4}, Vec2{-4, -6}},
		{Vec2{nan, -2}, Vec2{-3, -4}, Vec2{nan, -6}},
		{Vec2{pInf, -2}, Vec2{-3, -4}, Vec2{pInf, -6}},
		{Vec2{nInf, -2}, Vec2{-3, -4}, Vec2{nInf, -6}},
		{Vec2{pInf, -2}, Vec2{nInf, -4}, Vec2{nan, -6}},
	}
	for _, c := range cases {
		expected := c.result
		actual := c.A.Plus(c.B)
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.Plus(%v): expected: %v, got: %v", c.A, c.B, expected, actual)
		}
	}
}

func TestVec2Minus(t *testing.T) {
	cases := []struct {
		a, b, result Vec2
	}{
		{Vec2{0, 0}, Vec2{0, 0}, Vec2{0, 0}},
		{Vec2{1, 2}, Vec2{2, 6}, Vec2{-1, -4}},
		{Vec2{8, 9}, Vec2{3, 4}, Vec2{5, 5}},
		{Vec2{-1, -2}, Vec2{-3, -4}, Vec2{2, 2}},
		{Vec2{pInf, -2}, Vec2{-3, -4}, Vec2{pInf, 2}},
		{Vec2{nInf, -2}, Vec2{-3, -4}, Vec2{nInf, 2}},
		{Vec2{pInf, -2}, Vec2{nInf, -4}, Vec2{pInf, 2}},
		{Vec2{pInf, -2}, Vec2{pInf, -4}, Vec2{nan, 2}},
	}
	for _, c := range cases {
		expected := c.result
		actual := c.a.Minus(c.b)
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.Minus(%v): expected: %v, got: %v", c.a, c.b, expected, actual)
		}
	}
}

func TestVec2ScaledBy(t *testing.T) {
	cases := []struct {
		scalar    float32
		v, result Vec2
	}{
		{0, Vec2{0, 0}, Vec2{0, 0}},
		{0, Vec2{1, 2}, Vec2{0, 0}},
		{2, Vec2{1, 2}, Vec2{2, 4}},
		{-2, Vec2{1, 2}, Vec2{-2, -4}},
		{2, Vec2{-9, 2}, Vec2{-18, 4}},
		{0.001, Vec2{-9, 2}, Vec2{-0.009, 0.002}},
		{0.001, Vec2{pInf, 0}, Vec2{pInf, 0}},
		{0.001, Vec2{nInf, 0}, Vec2{nInf, 0}},
		{0.001, Vec2{nan, 0}, Vec2{nan, 0}},
		{0, Vec2{nInf, 0}, Vec2{nan, 0}},
	}
	for _, c := range cases {
		expected := c.result
		actual := c.v.ScaledBy(c.scalar)
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.Scaled(%v): expected: %v, got: %v", c.v, c.scalar, expected, actual)
		}
	}
}

func TestVec2Vec2Rand(t *testing.T) {
	cases := []Rect{
		Rect{},
		Rect{Vec2{10, 10}, Vec2{20, 20}},
		Rect{Vec2{-10, 20}, Vec2{50, 30}},
		Rect{Vec2{-0.001, 0}, Vec2{0.001, 0.0001}},
		Rect{Vec2{0, 0}, Vec2{10000, 100000}},
		Rect{Vec2{0, 0}, Vec2{1, pInf}},
		Rect{Vec2{nInf, 0}, Vec2{0, 0}},
		Rect{Vec2{nInf, 0}, Vec2{pInf, 0}},
		Rect{Vec2{nan, 0}, Vec2{1, 2}},
	}

	for _, rect := range cases {
		for i := 0; i < 4; i++ {
			vec := Vec2Rand(rect)

			if vec.X < rect.Min.X ||
				vec.X > rect.Max.X ||
				vec.Y < rect.Min.Y ||
				vec.Y > rect.Max.Y {
				t.Errorf("%v does not contain %v", rect, vec)
			}
		}
	}
}

func TestVec2PlusEquals(t *testing.T) {
	cases := []struct {
		A, B, result Vec2
	}{
		{Vec2{}, Vec2{}, Vec2{}},
		{Vec2{}, Vec2{1, 2}, Vec2{1, 2}},
		{Vec2{}, Vec2{-1, -2}, Vec2{-1, -2}},
		{Vec2{0.002, -9.32}, Vec2{0, 43.2}, Vec2{0.002, 33.88}},
		{Vec2{nan, pInf}, Vec2{0, 43.2}, Vec2{nan, pInf}},
	}

	for _, c := range cases {
		v := c.A
		v.PlusEquals(c.B)
		expected := c.result
		actual := v
		if !vec2Identical(expected, actual) {
			t.Errorf("%v.PlusEquals(%v): expected: %v, got: %v", c.A, c.B, expected, actual)
		}
	}
}

func TestVec2RandNormal(t *testing.T) {
	var sections [4]bool

	for i := 0; i < 100; i++ {
		v := Vec2RandNormal()

		length := v.X*v.X + v.Y*v.Y
		if !floatIdentical(length, 1) {
			t.Errorf("%v: expected length 1, got %v", v, length)
		}

		theta := math.Atan2(float64(v.Y), float64(v.X))
		switch {
		case theta > math.Pi*-1.0 && theta < math.Pi*-0.5:
			sections[0] = true
		case theta > math.Pi*-0.5 && theta < math.Pi*0:
			sections[1] = true
		case theta > math.Pi*0 && theta < math.Pi*0.5:
			sections[2] = true
		case theta > math.Pi*0.5 && theta < math.Pi*1:
			sections[3] = true
		}
	}

	for i := range sections {
		if sections[i] == false {
			t.Errorf("%d quarter of circle not covered", i)
		}
	}
}

func TestVec2Len2(t *testing.T) {
	cases := []struct {
		vec  Vec2
		len2 float32
	}{
		{Vec2{}, 0},
		{Vec2{1, 0}, 1},
		{Vec2{2, 0}, 4},
		{Vec2{0, 2}, 4},
		{Vec2{2, 2}, 8},
		{Vec2{-3, 4}, 25},
		{Vec2{3, 4}, 25},
	}

	for _, c := range cases {
		expected := c.len2
		actual := c.vec.Len2()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got %v", expected, actual)
		}
	}
}

func TestVec2Len(t *testing.T) {
	cases := []struct {
		vec    Vec2
		result float32
	}{
		{Vec2{}, 0},
		{Vec2{1, 0}, 1},
		{Vec2{2, 0}, 2},
		{Vec2{0, 2}, 2},
		{Vec2{-3, 4}, 5},
		{Vec2{3, 4}, 5},
		{Vec2{nan, 2}, nan},
		{Vec2{pInf, 2}, pInf},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.vec.Len()
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got %v", expected, actual)
		}
	}
}

func TestVec2Normal(t *testing.T) {
	cases := []struct {
		v, result Vec2
	}{
		{Vec2{}, Vec2{}},
		{Vec2{12, 0}, Vec2{1, 0}},
		{Vec2{0, -0.23}, Vec2{0, -1}},
		{
			Vec2{-12, -12},
			Vec2{float32(-math.Sqrt(0.5)), float32(-math.Sqrt(0.5))},
		},
		{
			Vec2{12, -12},
			Vec2{float32(math.Sqrt(0.5)), float32(-math.Sqrt(0.5))},
		},
		{Vec2{pInf, 2}, Vec2{1, 0}},
		{Vec2{-32, nInf}, Vec2{0, -1}},
		{Vec2{nan, nInf}, Vec2{0, 0}},
		{Vec2{pInf, pInf}, Vec2{0, 0}},
		{Vec2{nInf, pInf}, Vec2{0, 0}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.v.Normal()
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, got %v", expected, actual)
		}
	}
}

func TestVec2RotatedBy(t *testing.T) {
	cases := []struct {
		theta     float32
		v, result Vec2
	}{
		{0, Vec2{}, Vec2{}},
		{1.2, Vec2{}, Vec2{}},
		{0, Vec2{1, 2}, Vec2{1, 2}},
		{0, Vec2{-1, -2}, Vec2{-1, -2}},
		{math.Pi * 0.5, Vec2{1, 2}, Vec2{-2, 1}},
		{math.Pi * 1.0, Vec2{1, 2}, Vec2{-1, -2}},
		{math.Pi * 1.5, Vec2{1, 2}, Vec2{2, -1}},
		{math.Pi * 2.0, Vec2{1, 2}, Vec2{1, 2}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.v.RotatedBy(c.theta)
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestVec2Cross(t *testing.T) {
	cases := []struct {
		a, b   Vec2
		result float32
	}{
		{Vec2{}, Vec2{}, 0},
		{Vec2{1, 0}, Vec2{0, 1}, 1},
		{Vec2{2, 0}, Vec2{0, 1}, 2},
		{Vec2{1, 0}, Vec2{1, 0}, 0},
		{Vec2{-1, 0}, Vec2{1, 0}, 0},
		{Vec2{1, 0}, Vec2{1, 1}, 1},
		{Vec2{-1, 0}, Vec2{1, -1}, 1},
		{Vec2{4, 0}, Vec2{5, 5}, 20},
		{Vec2{4, -2}, Vec2{3, 7}, 34},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Cross(c.b)
		if !floatIdentical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}
