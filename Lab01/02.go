package main

import "fmt"

func main() {
	xs := [...]int{28, 3, 19, 8, 42, 5, 15}

	l := len(xs)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if xs[j] > xs[j+1] {
				tmp := xs[j]
				xs[j] = xs[j+1]
				xs[j+1] = tmp
			}
		}
	}

	fmt.Println("sortat: ", xs)
}
