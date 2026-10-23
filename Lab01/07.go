package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := make(map[string]int)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, word := range words {
			count[word]++
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("eroare la citire: ", err)
		return
	}

	for word, c := range count {
		fmt.Printf("%s: %d\n", word, c)
	}
}
