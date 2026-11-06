package main

import (
	"fmt"
	"math"
	"sync"
)

type Pos struct {
	x float64
	y float64
}

func dist(a Pos, b Pos) float64 {
	dx := a.x - b.x
	dy := a.y - b.y

	return math.Sqrt(dx*dx + dy*dy)
}

func robot(myIndex int, myPos Pos, newPos *Pos, chs *[][]chan Pos, adj *[][]bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for neighborIndex, isNeighbor := range (*adj)[myIndex] {
		if isNeighbor {
			(*chs)[myIndex][neighborIndex] <- myPos
		}
	}

	sumX := float64(0)
	sumY := float64(0)
	count := 1

	for fromIndex, line := range *adj {
		if line[myIndex] {
			otherPos := <-(*chs)[fromIndex][myIndex]
			sumX += otherPos.x
			sumY += otherPos.y
			count++
		}
	}

	*newPos = Pos{sumX / float64(count), sumY / float64(count)}
}

func readPositions() []Pos {
	var n int
	fmt.Scan(&n)

	pos := make([]Pos, n)
	for i := range pos {
		fmt.Scan(&pos[i].x, &pos[i].y)
	}

	return pos
}

func graph(positions []Pos) [][]bool {
	n := len(positions)

	adj := make([][]bool, n)
	for i := range n {
		adj[i] = make([]bool, n)

		for j := range n {
			if i != j && dist(positions[i], positions[j]) < 8.0 {
				adj[i][j] = true
			}
		}
	}

	return adj
}

func iteration(positions []Pos) []Pos {
	n := len(positions)
	adj := graph(positions)

	chs := make([][]chan Pos, n)
	for i := range n {
		chs[i] = make([]chan Pos, n)

		for j := range n {
			if adj[i][j] {
				chs[i][j] = make(chan Pos, 1)
			}
		}
	}

	newPositions := make([]Pos, n)

	wg := sync.WaitGroup{}
	for i, pos := range positions {
		wg.Add(1)
		go robot(i, pos, &newPositions[i], &chs, &adj, &wg)
	}

	wg.Wait()
	return newPositions
}

func converged(positions []Pos) bool {
	for i := range positions {
		for j := range positions {
			if dist(positions[i], positions[j]) >= 0.0001 {
				return false
			}
		}
	}

	return true
}

func main() {
	positions := readPositions()

	fmt.Println("initial:", positions)

	i := 0
	for !converged(positions) && i < 1000 {
		positions = iteration(positions)
		i++
	}

	fmt.Println("iteratii:", i)
	fmt.Println("final:", positions)
	fmt.Println("converged:", converged(positions))
	fmt.Println("punct final:", positions[0])
}
