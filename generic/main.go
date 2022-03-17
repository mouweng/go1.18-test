package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumInts(m map[string]int) int {
	var s int
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

func main() {
	ints := map[string]int{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("非泛型计算结果，SumInts: %v, SumFloats: %v\n", SumInts(ints), SumFloats(floats))
	
	fmt.Printf("泛型计算结果，SumInts :%v, SumFloats: %v\n", Sum(ints), Sum(floats))
}

/*
[Running] go run "./main.go"
非泛型计算结果，SumInts: 46, SumFloats: 62.97
泛型计算结果，SumInts :46, SumFloats: 62.97
*/
