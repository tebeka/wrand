package wrand

import (
	"fmt"
	"testing"
)

var s string

func BenchmarkRand(b *testing.B) {
	rw, err := New(map[string]int{"A": 10, "B": 20, "C": 30}, nil)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = rw.Rand()
	}
}

func BenchmarkRandBig(b *testing.B) {
	rw := genRand(b, 10_000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = rw.Rand()
	}
}

func genRand(b *testing.B, size int) WRand[string] {
	weights := make(map[string]int)
	for i := 1; i <= size; i++ {
		v := fmt.Sprintf("val-%d", i)
		weights[v] = i
	}

	rw, err := New(weights, nil)
	if err != nil {
		b.Fatal(err)
	}

	return rw
}
