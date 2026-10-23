package main

import "fmt"

func main() {
	xs := [...]int{28, 3, 19, 8, 42, 5, 15}
	ys := []int{}

	for _, num := range xs {
		if num%2 == 0 {
			ys = append(ys, num)
		}
	}

	fmt.Println("pare: ", ys)
}
