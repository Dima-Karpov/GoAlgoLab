package someMath

import (
	"math"
	"sort"
	"testing"
)

func BenchmarkSomeMath(b *testing.B) {
	rad := 1.2345
	for i := 0; i < b.N; i++ {
		_ = SomeMath(rad)
	}
}

// Бенчмарк для функции sort.Ints
func BenchmarkSortInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []int{5, 3, 4, 1, 2, 9, 7, 6, 8, 0}
		sort.Ints(data)
	}
}

// Test for math.Sqrt func
func TestMathSqrt(t *testing.T) {
	got := math.Sqrt(9)
	want := 3.0
	if got != want {
		t.Errorf("got %.1f but want %.1f", got, want)
	}
}
