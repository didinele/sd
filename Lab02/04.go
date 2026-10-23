package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("munca din goroutine")
		time.Sleep(5000 * time.Millisecond)
		ch <- "done"
	}()

	<-ch
}
