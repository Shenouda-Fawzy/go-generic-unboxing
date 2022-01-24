package main

import (
	"fmt"
	"github/go-generic/demo1"
)


func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		demo1.SumInts(ints),
		demo1.SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		demo1.SumIntsOrFloats[string, int64](ints),
		demo1.SumIntsOrFloats[string, float64](floats))
}

