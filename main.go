package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	// Catch Ctrl-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
		}
	}()

	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// FLAGS
