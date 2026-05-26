//go:build js
// +build js

package pry

func readFile(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type browserHistory struct {
	Records []string
}

// NewHistory constructs browserHistory instance
func NewHistory() (*browserHistory, error) {
	_ = "STUB: not implemented"

	// FIXME:
	// when localStorage is full, can be return an error
	return nil, nil
}

// Load unmarshal localStorage data into history's records
func (bh *browserHistory) Load() error { _ = "STUB: not implemented"; return nil }

// nothing to unmarashal

// Save saves marshaled history's records into localStorage
func (bh browserHistory) Save() error { _ = "STUB: not implemented"; return nil }

// Len returns amount of records in history
func (bh browserHistory) Len() int { _ = "STUB: not implemented"; return 0 }

// Add appends record into history's records
func (bh *browserHistory) Add(record string) { _ = "STUB: not implemented"; return }
