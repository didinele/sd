package main

import "fmt"

func receive(ch <-chan string) string {
	return <-ch
}

func send(ch chan<- string, msg string) {
	ch <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go send(ch1, "salut")
	go func() {
		ch2 <- receive(ch1)
	}()

	fmt.Println(receive(ch2))
}
