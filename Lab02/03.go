package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "salut1"
		ch <- "salut2"
	}()

	msg1 := <-ch
	fmt.Println(msg1)

	msg2 := <-ch
	fmt.Println(msg2)
}
