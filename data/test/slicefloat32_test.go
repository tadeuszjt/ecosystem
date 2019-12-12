package dataTest

import (
	"testing"
	"github.com/tadeuszjt/data"
)

func sliceFloat32Identical(a, b data.SliceFloat32) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	
	return true
}

func TestSliceFloat32Identical(t *testing.T) {
	cases := []struct{
		a, b   data.SliceFloat32
		result bool
	}{
		{
			data.SliceFloat32{},
			data.SliceFloat32{},
			true,
		},
		{
			data.SliceFloat32{12},
			data.SliceFloat32{},
			false,
		},
		{
			data.SliceFloat32{1, 2, 3, 4},
			data.SliceFloat32{1, 2, 3, 4},
			true,
		},
		{
			data.SliceFloat32{1, 2, 3, 4},
			data.SliceFloat32{1, 2, 4, 4},
			false,
		},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := sliceFloat32Identical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
		
}

func TestSliceFloat32Len(t *testing.T) {
	cases := []struct{
		slice  data.SliceFloat32
		result int
	}{
		{data.SliceFloat32{}, 0},
		{data.SliceFloat32{1, 2, 3}, 3},
		{nil, 0},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := c.slice.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSliceFloat32Swap(t *testing.T) {
	cases := []struct{
		i, j   int
		slice  data.SliceFloat32
		result data.SliceFloat32
	}{
		{0, 0, data.SliceFloat32{1}, data.SliceFloat32{1}},
		{0, 1, data.SliceFloat32{1, 2}, data.SliceFloat32{2, 1}},
		{1, 1, data.SliceFloat32{1, 2}, data.SliceFloat32{1, 2}},
		{1, 2, data.SliceFloat32{1, 2, 3, 4}, data.SliceFloat32{1, 3, 2, 4}},
		{3, 0, data.SliceFloat32{1, 2, 3, 4}, data.SliceFloat32{4, 2, 3, 1}},
	}
	
	for _, c := range cases {
		expected := c.result
		c.slice.Swap(c.i, c.j)
		actual := c.slice
		if !sliceFloat32Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestSliceFloat32Delete(t *testing.T) {
	cases := []struct{
		i      int
		slice  data.SliceFloat32
		result data.SliceFloat32
	}{
		{0, data.SliceFloat32{1}, data.SliceFloat32{}},
		{1, data.SliceFloat32{1, 2, 3}, data.SliceFloat32{1, 3}},
		{1, data.SliceFloat32{1, 2, 3, 4}, data.SliceFloat32{1, 4, 3}},
		{3, data.SliceFloat32{1, 2, 3, 4}, data.SliceFloat32{1, 2, 3}},
		{0, data.SliceFloat32{1, 2, 3, 4}, data.SliceFloat32{4, 2, 3}},
	}
	
	for _, c := range cases {
		expected := c.result
		c.slice.Delete(c.i)
		actual := c.slice
		if !sliceFloat32Identical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

