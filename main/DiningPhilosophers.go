package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	count sync.Mutex
}

type Philosopher struct {
	number           int
	host             *sync.Mutex
	eatingPhiloCount *int
	leftCS, rightCS  *ChopS
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; {
		p.host.Lock()
		if *p.eatingPhiloCount < 2 {

			p.leftCS.count.Lock()
			p.rightCS.count.Lock()

			fmt.Printf("starting to eat%d\n", p.number)
			*p.eatingPhiloCount = *p.eatingPhiloCount + 1
			p.host.Unlock()

			p.host.Lock()
			fmt.Printf("finishing eating%d\n", p.number)
			*p.eatingPhiloCount = *p.eatingPhiloCount - 1
			p.host.Unlock()

			p.rightCS.count.Unlock()
			p.leftCS.count.Unlock()
			i++

		} else {
			p.host.Unlock()
		}
	}
	wg.Done()
}

func main() {

	ResCount := 5
	var wg sync.WaitGroup
	eatingPhiloCount := 0
	var host sync.Mutex

	CSticks := make([]*ChopS, ResCount)
	for i := 0; i < ResCount; i++ {
		CSticks[i] = new(ChopS)
	}

	philosophers := make([]*Philosopher, ResCount)
	for i := 0; i < ResCount; i++ {
		philosophers[i] = &Philosopher{i + 1, &host, &eatingPhiloCount,
			CSticks[i], CSticks[(i+1)%ResCount]}
	}

	wg.Add(ResCount)

	for i := 0; i < ResCount; i++ {
		go philosophers[i].eat(&wg)
	}

	wg.Wait()

}
