package geomTest

import (
	. "github.com/tadeuszjt/geom/64"
	"testing"
)

func mat3Identical(a, b Mat3) bool {
	for i := range a {
		if !floatIdentical(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestMat3Identical(t *testing.T) {
	cases := []struct {
		a, b   Mat3
		result bool
	}{
		{Mat3Identity(), Mat3Identity(), true},
		{
			Mat3{0, 1, 2, 3, 4, 5, 6, 7, 8},
			Mat3{0, 1, 2, 3, 4, 5, 6, 7, 8},
			true,
		},
		{
			Mat3{0, 1, 2, 3, 4, 5, 6, 7, 8.1},
			Mat3{0, 1, 2, 3, 4, 5, 6, 7, 8},
			false,
		},
		{
			Mat3{pInf, 1, 2, 3, 4, 5, 6, 7, 8},
			Mat3{pInf, 1, 2, 3, 4, 5, 6, 7, 8},
			true,
		},
		{
			Mat3{pInf, 1, 2, 3, 4, 5, 6, 7, 8},
			Mat3{nInf, 1, 2, 3, 4, 5, 6, 7, 8},
			false,
		},
		{
			Mat3{pInf, 1, 2, 3, nan, 5, 6, 7, 8},
			Mat3{pInf, 1, 2, 3, nan, 5, 6, 7, 8},
			true,
		},
	}

	for _, c := range cases {
		expected := c.result
		actual := mat3Identical(c.a, c.b)

		if expected != actual {
			t.Errorf("a: %v, b: %v, expected: %v, got: %v",
				c.a, c.b, expected, actual)
		}
	}
}

func TestMat3Identity(t *testing.T) {
	expected := Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
	actual := Mat3Identity()
	if !mat3Identical(expected, actual) {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func TestMat3TimesVec2(t *testing.T) {
	for _, c := range []struct {
		mat    Mat3
		vec    Vec2
		bias   float64
		result Vec3
	}{
		{
			Mat3Identity(),
			Vec2{1, 1},
			1,
			Vec3{1, 1, 1},
		},
		{
			Mat3{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Vec2{10, 11},
			1,
			Vec3{35, 101, 167},
		},
		{
			Mat3{
				-3, pInf, 2.2,
				0, -38, 7,
				nan, 8, -0.1,
			},
			Vec2{-1, -2},
			-3,
			Vec3{nInf, 55, nan},
		},
		{
			Mat3{
				pInf, 0, 0,
				nInf, 0, 0,
				0.001, -0.002, 0.003,
			},
			Vec2{0, 1},
			2,
			Vec3{nan, nan, 0.004},
		},
	} {
		expected := c.result
		actual := c.mat.TimesVec2(c.vec, c.bias)
		if !vec3Identical(expected, actual) {
			t.Errorf("expected: %v, got: %v", expected, actual)
		}
	}
}

func TestMat3Camera2D(t *testing.T) {
	camera := Rect{
		Min: Vec2{10, 16},
		Max: Vec2{50, 32},
	}
	display := Rect{
		Min: Vec2{-1, -2},
		Max: Vec2{3, 4},
	}
	mat := Mat3Camera2D(camera, display)

	cases := []struct {
		point, result Vec2
	}{
		{Vec2{30, 24}, Vec2{1, 1}},
		{Vec2{10, 16}, Vec2{-1, -2}},
		{Vec2{50, 16}, Vec2{3, -2}},
		{Vec2{50, 32}, Vec2{3, 4}},
		{Vec2{10, 32}, Vec2{-1, 4}},
	}

	for _, c := range cases {
		actual := mat.TimesVec2(c.point, 1).Vec2()
		expected := c.result
		if !vec2Identical(expected, actual) {
			t.Errorf("point: %v: expected: %v, got: %v", c.point, expected, actual)
		}
	}
}

func TestMat3Times(t *testing.T) {
	cases := []struct {
		a, b, result Mat3
	}{
		{
			Mat3Identity(),
			Mat3Identity(),
			Mat3Identity(),
		},
		{
			Mat3{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Mat3{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Mat3{
				30, 36, 42,
				66, 81, 96,
				102, 126, 150,
			},
		},
		{
			Mat3{
				1, 2, 3,
				4, 5, 6,
				7, 8, 9,
			},
			Mat3{
				.1, .2, .3,
				.4, .5, .6,
				.7, .8, .9,
			},
			Mat3{
				3.0, 3.6, 4.2,
				6.6, 8.1, 9.6,
				10.2, 12.6, 15.0,
			},
		},
	}

	for _, c := range cases {
		actual := c.a.Product(c.b)
		expected := c.result
		if !mat3Identical(expected, actual) {
			t.Errorf("a: %v Times b: %v, expected: %v, got: %v",
				c.a, c.b, expected, actual)
		}
	}
}
