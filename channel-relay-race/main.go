package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	//first runner to his mark
	go Runner(baton)

	//start the race
	baton <- 1

	// wait for the race to finish
	wg.Wait()
}

func Runner(baton chan int) {
	var newRuner int

	// wait to recieve the baton

	runner := <-baton

	//start running around the track.
	fmt.Printf("Runner %d running With baton\n", runner)

	// new runner to the line
	if runner != 4 {
		newRuner = runner + 1
		fmt.Printf("Runner %d to the line.\n", newRuner)
		go Runner(baton)
	}

	//runing around the track.
	time.Sleep(100 * time.Millisecond)

	// is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// exchange the batonn for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRuner)

	baton <- newRuner
}
