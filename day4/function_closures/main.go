package main

import (
	"fmt"
	"math"
)

func mult(by int) func(int) int {
	return func(x int) int {
		return x * by
	}
}

func print_closure() {

	for i := 0; i <= 10; i++ {
		by_number := mult(i)
		for k := 0; k <= 10; k++ {
			result := by_number(k)
			fmt.Println(result)
		}
	}
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) set_x(x float64) {
	v.X = x
}

func main() {
	v := Vertex{3, 4}
	v.set_x(10)
	fmt.Println(v)
}
