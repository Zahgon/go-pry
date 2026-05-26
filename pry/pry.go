package pry

import (
	"io"
)

// Pry does nothing. It only exists so running code without go-pry doesn't throw an error.
func Pry(v ...interface{}) {
	_ = "STUB: not implemented"

	// Apply drops into a pry shell in the location required.
	return
}

func Apply(scope *Scope) { _ = "STUB: not implemented"; return }

type genericTTY interface {
	ReadRune() (rune, error)
	Size() (int, int, error)
	Close() error
}

func apply(
	scope *Scope,
	out io.Writer,
	tty genericTTY,
	filePath, filePathRaw string,
	lineNum int,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Down

// Up

// Right

// Left

// DELETE

// Backspace

// ? This happens on key press
//TAB
//ENTER

// Ctrl-D

func displayFilePosition(
	out io.Writer, filePathRaw, filePath string, lineNum int,
) {
	_ = "STUB: not implemented"
	return
}

// displaySuggestions renders the live autocomplete from GoCode.
func displaySuggestions(
	scope *Scope,
	out io.Writer,
	tty genericTTY,
	line string,
	index, promptWidth int,
) {
	_ = "STUB: not implemented"
	return
}
