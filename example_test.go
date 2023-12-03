package wrand_test

import (
	"fmt"

	"github.com/tebeka/wrand"
)

func ExampleWRand() {
	weights := map[string]int{
		"A": 20,
		"B": 30,
		"C": 50,
	}

	rw, err := wrand.New(weights, nil)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	for i := 0; i < 5; i++ {
		// We don't print here since the algorithm is not deterministic - flaky test.
		out := rw.Rand()
		if _, ok := weights[out]; !ok {
			fmt.Printf("error: unknown value - %q", out)
		}
	}

	// Output:
}
