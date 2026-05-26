package pry

// InterpretError is an error returned by the interpreter and shouldn't be
// passed to the user or running code.
type InterpretError struct {
	err error
}

func (a *InterpretError) Error() error { _ = "STUB: not implemented"; return nil }

// Append is a runtime replacement for the append function
func Append(arr interface{}, elems ...interface{}) (interface{}, *InterpretError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make is a runtime replacement for the make function
func Make(t interface{}, args ...interface{}) (interface{}, *InterpretError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close is a runtime replacement for the "close" function.
func Close(t interface{}) (interface{}, *InterpretError) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Len is a runtime replacement for the len function
func Len(t interface{}) (interface{}, *InterpretError) { _ = "STUB: not implemented"; return nil, nil }
