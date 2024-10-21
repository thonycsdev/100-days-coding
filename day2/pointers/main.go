package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var p *int
	i := 42
	p = &i
	*p = 30

	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)
	vertex := Vertex{1, 2}

	pv := &vertex
	fmt.Println(pv.X)    // mesma coisa
	fmt.Println((*pv).Y) // mesma coisa

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p2 = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p2, v2, v3)
}
