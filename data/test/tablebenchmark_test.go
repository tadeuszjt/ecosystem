package dataTest

import (
	"testing"
	. "github.com/tadeuszjt/data"
)

func BenchmarkTableLen(b *testing.B) {
	table := Table{
		&SliceInt{1, 2, 3, 4},
	}
	
	for i := 0; i < b.N; i++ {
		table.Len()
	}
}

func BenchmarkTableSwap(b *testing.B) {
	table := Table{
		&SliceInt{1, 2, 3, 4, 5, 6},
	}
	
	for i := 0; i < b.N; i++ {
		table.Swap(0, 3)
	}
}

func BenchmarkTableDelete(b *testing.B) {
	slice := SliceInt{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice2 SliceInt = slice[:]
	table := Table{ &slice2 }
	
	for i := 0; i < b.N; i++ {
		slice2 = slice[:]
		table.Delete(0)
	}
}

