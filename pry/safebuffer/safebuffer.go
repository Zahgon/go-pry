// safebuffer is a goroutine safe bytes.Buffer.
// From https://gist.github.com/arkan/5924e155dbb4254b64614069ba0afd81
package safebuffer

import (
	"bytes"
	"sync"
)

// Buffer is a goroutine safe bytes.Buffer
type Buffer struct {
	buffer bytes.Buffer
	mutex  sync.Mutex
}

// Write appends the contents of p to the buffer, growing the buffer as needed.
// It returns
// the number of bytes written.
func (s *Buffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// String returns the contents of the unread portion of the buffer
// as a string.  If the Buffer is a nil pointer, it returns "<nil>".
func (s *Buffer) String() string { _ = "STUB: not implemented"; return "" }
