package main

import "fmt"

func main() {
	fmt.Printf("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v\n", i)
	}

	fmt.Printf("Done")
}
