package pry

import (
	"go/token"

	"github.com/pkg/errors"
)

// ErrChanRecvFailed occurs when a channel is closed.
var ErrChanRecvFailed = errors.New("receive failed: channel closed")

// ErrChanRecvInSelect is an internal error that is used to indicate it's in a
// select statement.
var ErrChanRecvInSelect = errors.New("receive failed: in select")

var ErrDivisionByZero = errors.New("division by zero")

// DeAssign takes a *_ASSIGN token and returns the corresponding * token.
func DeAssign(tok token.Token) token.Token { _ = "STUB: not implemented"; return *new(token.Token) }

// ComputeBinaryOp executes the corresponding binary operation (+, -, etc) on two interfaces.
func ComputeBinaryOp(xI, yI interface{}, op token.Token) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Bool

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Num, uint

// Anything

// ComputeUnaryOp computes the corresponding unary (+x, -x) operation on an interface.
func (scope *Scope) ComputeUnaryOp(xI interface{}, op token.Token) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
