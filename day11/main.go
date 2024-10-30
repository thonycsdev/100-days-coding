package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test-data.txt")
	errorHandler(err)
	scanner := bufio.NewScanner(file)
	var names []string
	for scanner.Scan() {
		name := scanner.Text()
		names = append(names, name)
	}

	result := linear_search(names, "Ria Lloyd")

	if result == nil {
		fmt.Println("Name not found")
		return
	}
	name := names[*result]
	fmt.Println(name)

}

func linear_search(names []string, name string) *int {
	for idx, val := range names {
		if val == name {
			return &idx
		}
	}
	return nil

}

func errorHandler(err error) {
	if err != nil {
		log.Panic(err)
	}
}
