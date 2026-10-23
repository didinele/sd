package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	xs := []float32{28.1, 3.2, 19.3, 8.4, 42.5, 5.6, 15.7}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		fmt.Println("eroare la citire: ", err)
		return
	}

	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("failed string -> int")
		return
	}

	if i < 0 || i >= len(xs) {
		fmt.Println("bad index")
		return
	}

	xs = append(xs[:i], xs[i+1:]...)
	fmt.Println("dupa eliminare: ", xs)
}
