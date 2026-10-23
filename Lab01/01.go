package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func prim(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		fmt.Println("eroare la citire: ", err)
		return
	}

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("failed string -> int")
		return
	}

	if prim(n) {
		fmt.Printf("%d este prim\n", n)
	} else {
		fmt.Printf("%d nu este prim\n", n)
	}
}
