package main

import (
	"fmt"
	"sync"
)

func node(index int, vote int, adj *[][]bool, chs *[][]chan float64, reportCh chan<- float64, controlCh <-chan bool, finalValue *float64, wg *sync.WaitGroup) {
	defer wg.Done()

	x := float64(vote)

	inDegree := 0
	for _, line := range *adj {
		if line[index] {
			inDegree++
		}
	}

	for {
		for neighborIndex, isNeighbor := range (*adj)[index] {
			if isNeighbor {
				(*chs)[index][neighborIndex] <- x
			}
		}

		sum := x
		for fromIndex, line := range *adj {
			if line[index] {
				sum += <-(*chs)[fromIndex][index]
			}
		}
		x = sum / float64(inDegree+1)

		reportCh <- x
		if !<-controlCh {
			break
		}
	}

	*finalValue = x
}

func FloodingMedie(votes []int, adj [][]bool) []float64 {
	n := len(votes)

	chs := make([][]chan float64, n)
	for i := 0; i < n; i++ {
		chs[i] = make([]chan float64, n)

		for j := 0; j < n; j++ {
			if adj[i][j] {
				chs[i][j] = make(chan float64, 1)
			}
		}
	}

	reportCh := make(chan float64)
	controlChs := make([]chan bool, n)

	for i := range n {
		controlChs[i] = make(chan bool)
	}

	finalValues := make([]float64, n)

	wg := sync.WaitGroup{}

	for i, vote := range votes {
		wg.Add(1)
		go node(i, vote, &adj, &chs, reportCh, controlChs[i], &finalValues[i], &wg)
	}

	rounds := 0
	for {
		rounds++

		x := <-reportCh
		minX := x
		maxX := x

		for range n - 1 {
			x = <-reportCh
			minX = min(minX, x)
			maxX = max(maxX, x)
		}

		cont := maxX-minX >= 0.001 && rounds < 1000
		for i := range n {
			controlChs[i] <- cont
		}

		if !cont {
			break
		}
	}

	wg.Wait()
	return finalValues
}

func main() {
	votes := []int{1, 0, 1, 1, 0}
	adj := [][]bool{
		{false, true, true, false, false}, // 0 -> 1, 0 -> 2
		{false, false, true, true, false}, // 1 -> 2, 1 -> 3
		{false, false, false, true, true}, // 2 -> 3, 2 -> 4
		{true, false, false, false, true}, // 3 -> 4, 3 -> 0
		{true, true, false, false, false}, // 4 -> 0, 4 -> 1
	}

	values := FloodingMedie(votes, adj)
	fmt.Println(values)

	for i, x := range values {
		fmt.Printf("nod %d: %t\n", i, x > 0.5)
	}
}
