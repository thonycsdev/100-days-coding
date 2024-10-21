package main

import (
	"fmt"
	"math"
)

func compute(fn func(int, int) int) int {

	return fn(2, 8)
}

func main() {
	square := func(x, y int) int {
		return int(math.Pow(float64(x), float64(y)))
	}

	sum := func(x, y int) int {
		return x + y
	}

	result := compute(square)
	result_sum := compute(sum)
	fmt.Println(result)
	fmt.Println(result_sum)
}
