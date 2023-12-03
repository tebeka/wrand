// Package wrand provide weighted random.
package wrand

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

// WRand is weighted random generator.
type WRand[T comparable] struct {
	values []T
	totals []int
	n      int
	r      *rand.Rand
}

// New returns WRand from weights. If r is nil, it'll create new rand.Rand from the current time.
func New[T comparable](weights map[T]int, r *rand.Rand) (WRand[T], error) {
	if len(weights) == 0 {
		return WRand[T]{}, fmt.Errorf("empty weights")
	}

	values := make([]T, len(weights))
	totals := make([]int, len(weights))
	total, i := 0, 0
	for v, w := range weights {
		if w <= 0 {
			return WRand[T]{}, fmt.Errorf("weight of %v (%d) is less than 1", v, w)
		}

		total += w
		values[i], totals[i] = v, total
		i++
	}

	if r == nil {
		r = rand.New(rand.NewSource(time.Now().Unix()))
	}

	rw := WRand[T]{
		values: values,
		totals: totals,
		n:      total,
		r:      r,
	}
	return rw, nil
}

// Rand returns random element from the weighted population.
func (rw WRand[T]) Rand() T {
	n := rw.r.Intn(rw.n)
	i, _ := slices.BinarySearch(rw.totals, n)
	// TODO: Should it differ if found?
	return rw.values[i]
}
