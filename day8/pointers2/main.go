package main

import "fmt"

func MutateString(str *string) {
	*str = "Palloma"

}

func main() {
	name := "Anthony"
	MutateString(&name)
	fmt.Println(name)
}
