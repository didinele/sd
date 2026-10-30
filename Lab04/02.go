package main

import (
	"fmt"
	"sync"
)

func node(id int, isLeader *bool, diameter int, wg *sync.WaitGroup, adj *[][]bool, chs *[][]chan int, idToIndex *map[int]int) {
	defer wg.Done()

	max_id := id

	index := (*idToIndex)[id]
	for range diameter {
		// me -> other
		for neighborIndex, isNeighbor := range (*adj)[index] {
			if isNeighbor {
				(*chs)[index][neighborIndex] <- max_id
			}
		}

		// other -> me
		for neighborIndex, line := range *adj {
			if line[index] {
				max_id = max(max_id, <-(*chs)[neighborIndex][index])
			}
		}
	}

	if max_id == id {
		*isLeader = true
	}
}

func calcDiameter(ids []int, adj [][]bool) int {
	diameter := 0

	for index := range len(adj) {
		dist := make([]int, len(adj))
		for i := range dist {
			dist[i] = -1
		}

		queue := []int{index}
		dist[index] = 0

		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]

			for u := range adj[v] {
				if adj[v][u] && dist[u] == -1 {
					dist[u] = dist[v] + 1
					queue = append(queue, u)
				}
			}
		}

		for _, d := range dist {
			if d > diameter {
				diameter = d
			}
		}
	}

	return diameter
}

func FloodingGeneralizat(ids []int, adj [][]bool) int {
	leaderBools := make(map[int]*bool)
	idToIndex := make(map[int]int)

	for index, id := range ids {
		idToIndex[id] = index
	}

	chs := make([][]chan int, len(adj))
	for i := 0; i < len(adj); i++ {
		chs[i] = make([]chan int, len(adj))

		for j := 0; j < len(adj); j++ {
			if adj[i][j] {
				chs[i][j] = make(chan int, 1)
			} else {
				chs[i][j] = nil
			}
		}
	}

	wg := sync.WaitGroup{}

	diameter := calcDiameter(ids, adj)
	for _, id := range ids {
		leaderBool := false
		leaderBools[id] = &leaderBool

		wg.Add(1)
		go node(id, &leaderBool, diameter, &wg, &adj, &chs, &idToIndex)
	}

	wg.Wait()
	var leader int
	for id, isLeader := range leaderBools {
		if *isLeader {
			leader = id
		}
	}

	return leader
}

func main() {
	ids := []int{87, 12, 45, 3, 61}
	adj := [][]bool{
		{false, true, false, false, false}, // 0 -> 1
		{false, false, true, false, false}, // 1 -> 2
		{true, false, false, true, false},  // 2 -> 0, 2 -> 3
		{false, true, false, false, true},  // 3 -> 1, 3 -> 4
		{true, false, false, false, false}, // 4 -> 0
	}

	leader := FloodingGeneralizat(ids, adj)
	fmt.Println(leader)
}
