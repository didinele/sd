package main

import "fmt"

func main() {
	xs := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	l := len(xs)
	for i := 0; i < l/2; i++ {
		tmp := xs[i]
		xs[i] = xs[l-1-i]
		xs[l-1-i] = tmp
	}

	fmt.Println(xs)
}
