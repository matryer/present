package main

import (
	"log"
	"runtime"
)

// START OMIT

var (
	concurrent    = runtime.NumCPU()
	semaphoreChan = make(chan struct{}, concurrent)
)

func doWork(item int) {
	// push to channel will block if it's full
	semaphoreChan <- struct{}{}
	defer func() {
		// read from semaphoreChan to release a slot
		<-semaphoreChan
	}()

	// TODO: do work
	log.Println(item)
}

func main() {
	for i := 0; i < 10000; i++ {
		go doWork(i)
	}
}

// END OMIT

type Item struct{}
