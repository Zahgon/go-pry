//go:build js
// +build js

package pry

import (
	"io"
	"log"
	"syscall/js"
)

var tty = newWASMTTY()

func newWASMTTY() *wasmTTY { _ = "STUB: not implemented"; return nil }

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.SetOutput(tty)
}

func openTTY() (io.Writer, genericTTY) {
	_ = "STUB: not implemented"
	return *new(io.Writer), *new(genericTTY)
}

type wasmTTY struct {
	term js.Value
	r    io.Reader
}

func (t *wasmTTY) Write(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (t *wasmTTY) ReadRune() (rune, error) { _ = "STUB: not implemented"; return 0, nil }

func (t *wasmTTY) Size() (int, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (t *wasmTTY) Close() error { _ = "STUB: not implemented"; return nil }
