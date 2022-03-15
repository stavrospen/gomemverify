package main

import (
	"fmt"
	"sync"
)

var rounds = 10
var thread_count = 16
var multiplier = 10
var meg = 1024 * 1024
var size = multiplier * meg

func check(c1 []byte, x int) {
	if x == multiplier {
		return
	}
	check(c1, x+1)
	for i := 0; i < meg; i++ {
		if c1[(x*meg)+i] != 0xff {
			fmt.Printf("oops\n")
			panic("oops")
		}
	}
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	c1 := make([]byte, size)

	for i := 0; i < size; i++ {
		c1[i] = 0xff
	}

	for round := 0; round < rounds; round++ {
		check(c1, 0)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Go Options:\n")
	fmt.Printf(" rounds: %d\n", rounds)
	fmt.Printf(" threads: %d\n", thread_count)
	fmt.Printf(" multiplier: %d\n", multiplier)
	fmt.Printf(" size: %d\n", size)

	for i := 0; i < thread_count; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
}
