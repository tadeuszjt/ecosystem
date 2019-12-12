package geomTest

import (
	. "github.com/tadeuszjt/geom/32"
	"math"
	"testing"
)

func ori2Identical(a, b Ori2) bool {
	return floatIdentical(a.X, b.X) &&
		floatIdentical(a.Y, b.Y) &&
		floatIdentical(a.Theta, b.Theta)
}

func TestOri2(t *testing.T) {
	expected := Ori2{0, 0, 0}
	actual := Ori2{}
	if !ori2Identical(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestVec2(t *testing.T) {
	cases := []struct {
		o Ori2
		v Vec2
	}{
		{Ori2{}, Vec2{}},
		{Ori2{1, 2, 3}, Vec2{1, 2}},
		{Ori2{.1, .2, .3}, Vec2{.1, .2}},
		{Ori2{-1, -2, 3}, Vec2{-1, -2}},
		{Ori2{nan, pInf, nInf}, Vec2{nan, pInf}},
	}

	for _, c := range cases {
		expected := c.v
		actual := c.o.Vec2()
		if !vec2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2PlusEquals(t *testing.T) {
	cases := []struct {
		v, b, result Ori2
	}{
		{Ori2{}, Ori2{}, Ori2{}},
		{Ori2{0, 0, 0}, Ori2{1, 2, 3}, Ori2{1, 2, 3}},
		{Ori2{0, 0, 0}, Ori2{-1, -2, -3}, Ori2{-1, -2, -3}},
		{Ori2{1, 2, 3}, Ori2{4, 5, 6}, Ori2{5, 7, 9}},
		{Ori2{1, 2, 3}, Ori2{-4, -5, -6}, Ori2{-3, -3, -3}},
		{Ori2{nan, 2, 3}, Ori2{4, 5, 6}, Ori2{nan, 7, 9}},
	}

	for _, c := range cases {
		expected := c.result
		c.v.PlusEquals(c.b)
		actual := c.v
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2ScaledBy(t *testing.T) {
	cases := []struct {
		scalar    float32
		o, result Ori2
	}{
		{0, Ori2{}, Ori2{}},
		{0, Ori2{1, 2, 3}, Ori2{0, 0, 0}},
		{1, Ori2{1, 2, 3}, Ori2{1, 2, 3}},
		{-2, Ori2{1, 2, 3}, Ori2{-2, -4, -6}},
		{nan, Ori2{1, 2, 3}, Ori2{nan, nan, nan}},
		{0.001, Ori2{1, 2, 3}, Ori2{0.001, 0.002, 0.003}},
		{pInf, Ori2{1, -2, 3}, Ori2{pInf, nInf, pInf}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.ScaledBy(c.scalar)
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Mat3Transform(t *testing.T) {
	cases := []struct {
		o      Ori2
		v      Vec2
		result Vec3
	}{
		{Ori2{}, Vec2{}, Vec3{0, 0, 1}},
		{Ori2{1, 2, 0}, Vec2{0, 0}, Vec3{1, 2, 1}},
		{Ori2{1, 2, 0}, Vec2{3, 4}, Vec3{4, 6, 1}},
		{Ori2{3, 4, math.Pi / 2}, Vec2{1, 2}, Vec3{1, 5, 1}},
		{Ori2{3, 4, -math.Pi / 2}, Vec2{1, 2}, Vec3{5, 3, 1}},
		{Ori2{-2, 8, math.Pi}, Vec2{3, -2}, Vec3{-5, 10, 1}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.Mat3Transform().TimesVec2(c.v, 1)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Vec2(t *testing.T) {
	cases := []struct {
		o      Ori2
		result Vec3
	}{
		{Ori2{}, Vec3{}},
		{Ori2{1, 2, 3}, Vec3{1, 2, 3}},
		{Ori2{-1, -2, -3}, Vec3{-1, -2, -3}},
		{Ori2{0.001, 0.002, 0.003}, Vec3{0.001, 0.002, 0.003}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.o.Vec3()
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestOri2Times(t *testing.T) {
	cases := []struct {
		a, b, result Ori2
	}{
		{Ori2{}, Ori2{}, Ori2{}},
		{Ori2{0, 0, 0}, Ori2{1, 2, 3}, Ori2{0, 0, 0}},
		{Ori2{1, 0.2, 3}, Ori2{0.4, 5, 0.6}, Ori2{0.4, 1, 1.8}},
		{Ori2{-1, -2, -3}, Ori2{4, 5, 6}, Ori2{-4, -10, -18}},
		{Ori2{nan, pInf, nInf}, Ori2{-4, -5, -6}, Ori2{nan, nInf, pInf}},
		{Ori2{nan, pInf, nInf}, Ori2{0, 0, 0}, Ori2{nan, nan, nan}},
	}

	for _, c := range cases {
		expected := c.result
		actual := c.a.Times(c.b)
		if !ori2Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
