package pry

// Package represents a Go package for use with pry
type Package struct {
	Name      string
	Functions map[string]interface{}
}

func (p Package) Keys() []string { _ = "STUB: not implemented"; return nil }

func (p Package) Get(key string) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }
