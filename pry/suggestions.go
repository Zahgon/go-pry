package pry

import (
	"regexp"
)

var suggestionsRegexp = regexp.MustCompile("[.0-9a-zA-Z]+$")

func (s *Scope) SuggestionsPry(line string, index int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keyser interface {
	Keys() []string
}

type getter interface {
	Get(string) (interface{}, bool)
}

func get(v interface{}, key string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func keys(v interface{}) []string { _ = "STUB: not implemented"; return nil }
