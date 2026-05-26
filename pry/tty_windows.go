//go:build windows
// +build windows

package pry

import (
	"io"
)

func openTTY() (io.Writer, genericTTY) {
	_ = "STUB: not implemented"
	return *new(io.Writer), *new(genericTTY)
}
