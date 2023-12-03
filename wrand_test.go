package wrand_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
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

	r := rand.New(rand.NewSource(9))
	rw, err := wrand.New(weights, r)
	require.NoError(t, err)

	counts := make(map[string]int)
	const size = 1_000_000
	for i := 0; i < size; i++ {
		counts[rw.Rand()]++
	}

	for s, c := range counts {
		actual := float64(c) / size
		wanted := float64(weights[s]) / total
		require.InDelta(t, wanted, actual, 0.05)
	}
}
