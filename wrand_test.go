package wrand_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/tebeka/wrand"
)

func TestWRand(t *testing.T) {
	// values add up to 100 to make it easier to calculate weights.
	weights := map[string]int{
		"A": 20,
		"B": 30,
		"C": 50,
	}

	total := 0.0
	for _, v := range weights {
		total += float64(v)
	}

	r := rand.New(rand.NewPCG(9, 0))
	rw, err := wrand.New(weights, r)
	if err != nil {
		t.Fatal(err)
	}

	counts := make(map[string]int)
	const size = 1_000_000
	for i := 0; i < size; i++ {
		counts[rw.Rand()]++
	}

	for s, c := range counts {
		actual := float64(c) / size
		wanted := float64(weights[s]) / total
		if math.Abs(wanted-actual) > 0.05 {
			t.Errorf("weight for %s: wanted %f, got %f (diff %f)", s, wanted, actual, math.Abs(wanted-actual))
		}
	}
}
