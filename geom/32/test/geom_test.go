package geomTest

import (
	"math"
	"testing"
)

var (
	nan  = float32(math.NaN())
	pInf = float32(math.Inf(0))
	nInf = float32(math.Inf(-1))
)

func floatIdentical(a32, b32 float32) bool {
	a := float64(a32)
	b := float64(b32)
	return math.IsNaN(a) && math.IsNaN(b) ||
		math.IsInf(a, -1) && math.IsInf(b, -1) ||
		math.IsInf(a, 1) && math.IsInf(b, 1) ||
		math.Abs(a-b) < 0.000001
}

func TestFloatIdentical(t *testing.T) {
	cases := []struct {
		a, b   float32
		result bool
	}{
		{0, 0, true},
		{0, 0.00001, false},
		{-123.45, -123.45, true},
		{-123.45, 123.45, false},
		{nan, 0, false},
		{nan, -nan, true},
		{nan, nan, true},
		{nan, pInf, false},
		{nan, nInf, false},
		{pInf, nInf, false},
		{-pInf, nInf, true},
		{pInf, pInf, true},
	}

	for _, c := range cases {
		expected := c.result
		actual := floatIdentical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
