package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
)

type Msg struct {
	Value int
	Done  bool // liderul anunta ca a fost ales
}

// realizeaza operatiile interne nodului
func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	out <- Msg{id, false}

	for {
		msg := <-in

		if msg.Done {
			if !*isLeader {
				out <- msg
			}
			break
		} else if msg.Value == id {
			*isLeader = true
			out <- Msg{id, true}
		} else if msg.Value > id {
			out <- msg
		}
	}
}

// returneaza id-ul liderului
func RingFlooding(ids []int) int {
	leaderBools := make(map[int]*bool)
	chs := make([]chan Msg, len(ids))

	for i := range ids {
		chs[i] = make(chan Msg, 1)
	}

	wg := sync.WaitGroup{}

	for i, id := range ids {
		leaderBool := false
		leaderBools[id] = &leaderBool

		outCh := chs[(i+1)%len(ids)]

		wg.Add(1)
		go node(id, chs[i], outCh, &leaderBool, &wg)
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
	n := rand.Intn(100)
	ids := make([]int, 0)

	for len(ids) < n {
		r := rand.Intn(1000)
		if !slices.Contains(ids, r) {
			ids = append(ids, r)
		}
	}

	fmt.Println(ids)

	leader := RingFlooding(ids)
	fmt.Println(leader)
}
