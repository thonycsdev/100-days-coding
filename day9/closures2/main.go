package main

import "fmt"

func multiply(number int) func(x int) int {
	return func(x int) int {
		return number * x
	}
}
func main() {
	for x := 0; x <= 10; x++ {
		for y := 0; y <= 10; y++ {
			fmt.Printf("%d x %d = %d\n", x, y, x*y)
		}
	}
}
