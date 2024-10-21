package main

import (
	"fmt"
	"math"
)

func main() {

	sum := 1
	for sum <= 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println("Hello")
	result := Sqrt(7)
	fmt.Println(result)

	defer fmt.Println("World")
}

func withWhile() {
	sum := 1
	for sum <= 1000 {
		sum += sum
	}

	fmt.Println(sum)

}

func infinite() {
	for {
	}
}

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim

}

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 1000; i++ {
		v := (z*z - x) / (2 * z)
		if v == z {
			return v
		}
		z -= v
	}
	return z
}
