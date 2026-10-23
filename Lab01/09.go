package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	xs := [...]int{2, 7, 11, 8, 4, 9}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		fmt.Println("eroare la citire: ", err)
		return
	}

	s, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("failed string -> int")
		return
	}

	seen := make(map[int]int)
	for j, x := range xs {
		i, ok := seen[s-x]
		if ok {
			fmt.Println(i, j)
		}
		seen[x] = j
	}
}
