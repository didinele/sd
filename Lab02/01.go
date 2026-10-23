package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "salut"
	}()

	msg := <-ch
	fmt.Println(msg)
}
