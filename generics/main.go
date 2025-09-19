package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"key1": 32,
		"key2": 64,
	}

	floats := map[string]float64{
		"key1": 32.01,
		"key2": 64.54,
	}

	fmt.Printf("Non-Generic Sums: %v and %.2f\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sum: %v and %.2f\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
