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
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		demo1.SumIntsOrFloats[string, int64](ints),
		demo1.SumIntsOrFloats[string, float64](floats))

	
	// You can omit the arguments because the compiler can infer them
	fmt.Printf("Generic Sums - with type args removed: %v and %v\n",
		demo1.SumIntsOrFloats(ints),
		demo1.SumIntsOrFloats(floats))
	
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		demo1.SumGeneric(ints),
		demo1.SumGeneric(floats))
}


// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
