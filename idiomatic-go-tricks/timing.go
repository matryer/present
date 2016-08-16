package main

import (
	"log"
	"time"
)

// START OMIT

func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(name, "took", d)
	}
}

func FunkyFunc() {
	stop := StartTimer("FunkyFunc")
	defer stop()

	time.Sleep(1 * time.Second)
}

// END OMIT
