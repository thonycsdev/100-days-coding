package main

import (
	"bufio"
	"fmt"
	"os"
)

func saveStringInFile() {

	file, _ := os.Create("string.txt")
	defer file.Close()

	name := "Anthony"
	bytesWritten, _ := file.WriteString(name)
	fmt.Println("Bytes Written: ", bytesWritten)
}

func saveLinesToFiles() {
	file, _ := os.Create("string.txt")
	defer file.Close()
	names := []string{"Palloma", "Anthony", "Chico", "Mingau"}
	for _, name := range names {
		bytesWritten, _ := fmt.Fprintln(file, name)
		fmt.Println("Bytes Written: ", bytesWritten)
	}
}

func readContentFile() {
	content, _ := os.ReadFile("test-data.txt")
	//read file abre todo o arquivo na memoria, nao importa o tamanho. CUIDADO !
	fmt.Println(string(content))

}

func readFileLineByLine() {
	content, _ := os.Open("string.txt")
	defer content.Close()

	scanner := bufio.NewScanner(content)
	scanner.Scan()
	fmt.Println(scanner.Text())
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// 	fmt.Println("=====")
	// }
}

func main() {
	saveLinesToFiles()
	readFileLineByLine()
}
