package dataTest

import (
	"testing"
	. "github.com/tadeuszjt/data"
)

func tableIdentical(a, b Table) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		switch sa := a[i].(type) {
			case *SliceInt:
				sb, ok := b[i].(*SliceInt)
				if !ok || !sliceIntIdentical(*sa, *sb) {
					return false
				}
				
			case *SliceFloat32:
				sb, ok := b[i].(*SliceFloat32)
				if !ok || !sliceFloat32Identical(*sa, *sb) {
					return false
				}
				
			default:
				panic("testSliceIntIdentical: unrecognised slice type")
		}
	}
	
	return true
}

func TestTableIdentical(t *testing.T) {
	cases := []struct{
		a, b   Table
		result bool
	}{
		{
			Table{},
			Table{},
			true,
		},
		{
			Table{},
			Table{ &SliceInt{} },
			false,
		},
		{
			Table{ &SliceInt{} },
			Table{ &SliceInt{} },
			true,
		},
		{
			Table{ &SliceFloat32{} },
			Table{ &SliceInt{} },
			false,
		},
		{
			Table{ &SliceInt{1, 2, 3} },
			Table{ &SliceInt{1, 2, 3} },
			true,
		},
		{
			Table{ &SliceInt{1, 2, 3} },
			Table{ &SliceInt{1, 2, 4} },
			false,
		},
		{
			Table{
				&SliceInt{1, 2, 3},
				&SliceFloat32{1, 2, 3},
			},
			Table{
				&SliceInt{1, 2, 3},
				&SliceFloat32{1, 2, 3},
			},
			true,
		},
		{
			Table{
				&SliceInt{1, 2, 3},
				&SliceFloat32{1, 2, 3},
			},
			Table{
				&SliceInt{1, 2, 3},
				&SliceFloat32{1, 2, 3.1},
			},
			false,
		},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := tableIdentical(c.a, c.b)
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestTableLen(t *testing.T) {
	cases := []struct{
		table  Table
		result int
	}{
		{
			Table{ &SliceInt{} },
			0,
		},
		{
			Table{ &SliceInt{1, 2, 3} },
			3,
		},
		{
			Table{
				&SliceInt{1, 2, 3},
				&SliceFloat32{1, 2, 3},
			},
			3,
		},
	}
	
	for _, c := range cases {
		expected := c.result
		actual := c.table.Len()
		if expected != actual {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestTableSwap(t *testing.T) {
	cases := []struct{
		i, j   int
		table  Table
		result Table
	}{
		{
			0, 0,
			Table{},
			Table{},
		},
		{
			0, 0,
			Table{ &SliceInt{1} },
			Table{ &SliceInt{1} },
		},
		{
			1, 3,
			Table{ &SliceInt{1, 2, 3, 4} },
			Table{ &SliceInt{1, 4, 3, 2} },
		},
		{
			2, 0,
			Table{
				&SliceInt{1, 2, 3, 4},
				&SliceFloat32{.1, .2, .3, .4},
			},
			Table{
				&SliceInt{3, 2, 1, 4},
				&SliceFloat32{.3, .2, .1, .4},
			},
		},
	}
	
	for _, c := range cases {
		expected := c.result
		c.table.Swap(c.i, c.j)
		actual := c.table
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}

func TestTableDelete(t *testing.T) {
	cases := []struct{
		i      int
		table  Table
		result Table
	}{
		{
			0,
			Table{},
			Table{},
		},
		{
			0,
			Table{
				&SliceInt{1, 2, 3, 4},
				&SliceFloat32{1, 2, 3, 4},
			},
			Table{
				&SliceInt{4, 2, 3},
				&SliceFloat32{4, 2, 3},
			},
		},
		{
			1,
			Table{
				&SliceInt{1, 2, 3, 4},
				&SliceFloat32{1, 2, 3, 4},
			},
			Table{
				&SliceInt{1, 4, 3},
				&SliceFloat32{1, 4, 3},
			},
		},
		{
			3,
			Table{
				&SliceInt{1, 2, 3, 4},
				&SliceFloat32{1, 2, 3, 4},
			},
			Table{
				&SliceInt{1, 2, 3},
				&SliceFloat32{1, 2, 3},
			},
		},
	}
	
	for _, c := range cases {
		expected := c.result
		c.table.Delete(c.i)
		actual := c.table
		if !tableIdentical(expected, actual) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
