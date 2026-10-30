package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() {
	l.counter++
}

func (l *LamportClock) Send() {
	l.Tick()
}

func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.Tick()
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string // "internal", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a Event, b Event) bool {
	return a.internal_counter < b.internal_counter
}

func rutina(nextCh chan<- Message, prevCh <-chan Message, doneCh chan<- []Event, id int) {
	clock := LamportClock{}
	events := []Event{}

	// simulam viteze diferite ale ceasurilor
	clock.counter += rand.Intn(7)
	clock.Tick()
	events = append(events, Event{id, clock.counter, "internal"})

	clock.counter += rand.Intn(7)
	clock.Send()
	nextCh <- Message{id, clock.counter}
	events = append(events, Event{id, clock.counter, "send"})

	clock.counter += rand.Intn(7)
	msg := <-prevCh
	clock.Receive(msg.ts)
	events = append(events, Event{id, clock.counter, "recv"})

	doneCh <- events
}

func main() {
	doneCh := make(chan []Event)

	oneToTwoCh := make(chan Message, 1)
	twoToThreeCh := make(chan Message, 1)
	threeToOneCh := make(chan Message, 1)

	go rutina(oneToTwoCh, threeToOneCh, doneCh, 1)
	go rutina(twoToThreeCh, oneToTwoCh, doneCh, 2)
	go rutina(threeToOneCh, twoToThreeCh, doneCh, 3)

	all := []Event{}
	all = append(all, <-doneCh...)
	all = append(all, <-doneCh...)
	all = append(all, <-doneCh...)

	sort.Slice(all, func(i int, j int) bool {
		return less(all[i], all[j])
	})

	for _, e := range all {
		fmt.Printf("P%d  %3d  %s\n", e.process, e.internal_counter, e.kind)
	}
}
