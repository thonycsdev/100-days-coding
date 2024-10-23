package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	find_duplicates_from_terminal()
}

func find_duplicates_from_terminal() {
	file_path := os.Args[1:]
	fmt.Println(file_path)
	file, err := os.Open(file_path[0])
	if err != nil {
		log.Panicln(err)
	}
	scanner := bufio.NewScanner(file)
	line_number := hash_file_to_map(scanner)
	print_line_number(&line_number)

}

func find_duplicates_1() {
	input := bufio.NewScanner(os.Stdin)
	line_number := hash_file_to_map(input)
	print_line_number(&line_number)

}

func hash_file_to_map(file *bufio.Scanner) map[string]int {
	line_number := make(map[string]int)
	for file.Scan() {
		key := file.Text()
		if _, ok := line_number[key]; ok {
			line_number[key]++
		} else {
			line_number[key] = 1
		}
	}

	return line_number

}

// print map de string int
func print_line_number(line_number *map[string]int) {

	for key, val := range *line_number {
		if val > 1 {
			fmt.Printf("%s. Amount: %d \n", key, val)
		}
	}

}
