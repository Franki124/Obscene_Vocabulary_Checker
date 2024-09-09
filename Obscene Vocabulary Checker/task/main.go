package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var filename string

	fmt.Scanln(&filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordMap := make(map[string]bool)

	for scanner.Scan() {
		word := scanner.Text()
		wordMap[strings.ToLower(word)] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(os.Stdin)

	for inputScanner.Scan() {
		inputWord := inputScanner.Text()
		if inputWord == "exit" {
			fmt.Println("Bye!")
			break
		}

		lowerWord := strings.ToLower(inputWord)
		if _, exists := wordMap[lowerWord]; exists {
			fmt.Println(strings.Repeat("*", len(inputWord)))
		} else {
			fmt.Println(inputWord)
		}
	}

}
