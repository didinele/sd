package main

import "fmt"

func main() {
	h := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	names := []string{}
	for name := range h {
		names = append(names, name)
	}

	l := len(names)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if h[names[j]] < h[names[j+1]] {
				tmp := names[j]
				names[j] = names[j+1]
				names[j+1] = tmp
			}
		}
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
