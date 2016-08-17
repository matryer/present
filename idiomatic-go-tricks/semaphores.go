package main

import (
	"log"
	"time"
)

// START OMIT

var (
	concurrent    = 5
	semaphoreChan = make(chan struct{}, concurrent)
)

func doWork(item int) {
	semaphoreChan <- struct{}{} // block while full
	go func() {
		defer func() {
			<-semaphoreChan // read to release a slot
		}()
		log.Println(item)
		time.Sleep(500 * time.Millisecond)
	}()
}

func main() {
	for i := 0; i < 10000; i++ {
		doWork(i)
	}
	time.Sleep(5 * time.Second)
}

// END OMIT
