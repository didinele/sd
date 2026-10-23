package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// grid 100x100. tinta pentru roboti e (0,0)
type Pos struct {
	x int
	y int
}

func randomStartPos() Pos {
	x := rand.Int() % 100
	y := rand.Int() % 100

	return Pos{x, y}
}

func nextPos(current Pos) Pos {
	ax, ay := current.x, current.y

	if ax < 0 {
		ax = -ax
	}

	if ay < 0 {
		ay = -ay
	}

	if current.x != 0 {
		if current.x > 0 {
			current.x--
		} else {
			current.x++
		}
	} else if current.y != 0 {
		if current.y > 0 {
			current.y--
		} else {
			current.y++
		}
	}

	return current
}

func robot(toOther chan<- Pos, fromOther <-chan Pos, toDone chan<- int, robotNum int) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for p := range fromOther {
			fmt.Printf("robotul %d a primit: %+v\n", robotNum, p)
		}
	}()

	currentLocation := randomStartPos()
	toOther <- currentLocation

	for currentLocation.x != 0 || currentLocation.y != 0 {
		currentLocation = nextPos(currentLocation)
		toOther <- currentLocation
	}
	close(toOther)

	wg.Wait()
	toDone <- robotNum
}

func main() {
	r1r2ch := make(chan Pos)
	r2r1ch := make(chan Pos)

	rDoneCh := make(chan int)

	go robot(r1r2ch, r2r1ch, rDoneCh, 1)
	go robot(r2r1ch, r1r2ch, rDoneCh, 2)

	fmt.Printf("robotul %d a terminat\n", <-rDoneCh)
	fmt.Printf("robotul %d a terminat\n", <-rDoneCh)
}
