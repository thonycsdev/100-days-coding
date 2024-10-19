package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Pi is: ", math.Pi)
	fmt.Println("Sum: ", add(1, 1))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(10))
	fmt.Println(c, python, java)

}
