package main

import (
	"time"

	"github.com/d4l3k/go-pry/pry"
)

func prying() { _ = "STUB: not implemented"; return }

func main() {
	c := make(chan bool)
	go func() {
		prying()
		pry.Pry()
		c <- true
	}()
	<-c
	for {
		time.Sleep(time.Second)
	}
}
