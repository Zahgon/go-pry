//go:build !js
// +build !js

package pry

import (
	"io/ioutil"
)

var readFile = ioutil.ReadFile

var historyFile = ".go-pry_history"

type ioHistory struct {
	FileName string
	FilePath string
	Records  []string
}

// NewHistory constructs ioHistory instance
func NewHistory() (*ioHistory, error) { _ = "STUB: not implemented"; return nil, nil }

// Load unmarshal history file into history's records
func (h *ioHistory) Load() error { _ = "STUB: not implemented"; return nil }

// Save saves marshaled history's records into file
func (h ioHistory) Save() error { _ = "STUB: not implemented"; return nil }

// Len returns amount of records in history
func (h ioHistory) Len() int { _ = "STUB: not implemented"; return 0 }

// Add appends record into history's records
func (h *ioHistory) Add(record string) { _ = "STUB: not implemented"; return }
