package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robot(ch chan<- string) {
	d := time.Duration(rand.Int()%5000) * time.Millisecond
	time.Sleep(d)

	ch <- "done"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go robot(ch1)
	go robot(ch2)

	select {
	case <-ch1:
		fmt.Println("robotul 1 a ajuns primul")
	case <-ch2:
		fmt.Println("robotul 2 a ajuns primul")
	}
}
