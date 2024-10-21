package main

import "fmt"

func main() {
	printing_range()
}

func slice_references() {
	names := [4]string{ //array
		"Anthony",
		"Palloma",
		"Chico",
		"Mingau",
	}

	fmt.Println(names)

	a := names[0:2] //slices
	b := names[1:]  // slices

	b[0] = "XXXX"
	fmt.Println(a, b)
	fmt.Println(names)
	b = append(b, "Princesa")
	fmt.Println(len(names), cap(b))
	fmt.Println(a, b)
}

func nil_slices() {
	var nil_slice []int
	fmt.Println(nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("Nil!")
	}
}

func creating_slices_make() {
	a := make([]int, 5)    // cria um array com len(5) e uma slice para esse array a
	b := make([]int, 2, 5) // criar um array com len(2) e capacidade 5

	fmt.Println(a)
	fmt.Println(b)

}

func printing_range() {

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
