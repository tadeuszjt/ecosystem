package geomTest

import (
	"math"
	"testing"
)

var (
	nan  = math.NaN()
	pInf = math.Inf(0)
	nInf = math.Inf(-1)
)

func floatIdentical(a, b float64) bool {
	return math.IsNaN(a) && math.IsNaN(b) ||
		math.IsInf(a, -1) && math.IsInf(b, -1) ||
		math.IsInf(a, 1) && math.IsInf(b, 1) ||
		math.Abs(a-b) < 0.0000001
}

func TestFloatIdentical(t *testing.T) {
	cases := []struct {
		a, b   float64
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
