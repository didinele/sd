package main

import "fmt"

func main() {
	emp := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	dep := make(map[int][]string)
	for name, id := range emp {
		dep[id] = append(dep[id], name)
	}

	for id, names := range dep {
		fmt.Printf("Departamentul %d: %v\n", id, names)
	}
}
