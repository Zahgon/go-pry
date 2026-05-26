package main

import (
	"log"
	"regexp"
)

const out = "fuzz/corpus/"

var (
	exampleRegexpQuotes = regexp.MustCompile("(?s)InterpretString\\(`(.*?)`\\)")
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }
