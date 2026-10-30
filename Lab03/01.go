package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// TODO:
// Calculeaza decalajul theta si round -trip -ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.
		Millisecond)
	t2 := s_clock.Now()
	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.
		Millisecond)
	t3 := s_clock.Now()
	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.
		Millisecond)
	t4 := c_clock.Now()
	return estimate(t1, t2, t3, t4)
}

func main() {
	initialOffset := 150 * time.Millisecond

	s_clock := &Clock{0}

	c_clock := &Clock{initialOffset}
	fmt.Println("decalaj initial: ", c_clock.offset-s_clock.offset)

	theta, delta := measure(c_clock, s_clock)
	fmt.Println("theta: ", theta, " delta: ", delta)

	c_clock.offset += theta
	fmt.Println("decalaj dupa corectie: ", c_clock.offset-s_clock.offset)

	// b
	fmt.Println()
	fmt.Println("(b)")
	c_clock = &Clock{initialOffset}
	fmt.Println("decalaj initial: ", c_clock.offset-s_clock.offset)

	bestTheta, bestDelta := measure(c_clock, s_clock)
	minTheta := bestTheta
	fmt.Printf("1: theta: %v  delta: %v\n", bestTheta, bestDelta)

	for i := 2; i <= 10; i++ {
		theta, delta := measure(c_clock, s_clock)
		fmt.Printf("%d: theta: %v  delta: %v\n", i, theta, delta)

		if delta < bestDelta {
			bestDelta = delta
			bestTheta = theta
		}
		if theta < minTheta {
			minTheta = theta
		}
	}

	fmt.Println("delta minim: ", bestDelta, " theta corespunzator: ", bestTheta)
	fmt.Println("theta minim: ", minTheta)

	c_clock.offset += bestTheta
	fmt.Println("decalaj dupa corectie: ", c_clock.offset-s_clock.offset)
}
