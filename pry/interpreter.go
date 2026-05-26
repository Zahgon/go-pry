package pry

import (
	"go/ast"
	"go/token"
	"reflect"
	"sync"

	"github.com/pkg/errors"

	"go/types"
	// Used by types for import determination
)

var (
	// ErrChanSendFailed occurs when a channel is full or there are no receivers
	// available.
	ErrChanSendFailed = errors.New("failed to send, channel full or no receivers")

	// ErrBranchBreak is an internal error thrown when a for loop breaks.
	ErrBranchBreak = errors.New("branch break")
	// ErrBranchContinue is an internal error thrown when a for loop continues.
	ErrBranchContinue = errors.New("branch continue")
)

// Scope is a string-interface key-value pair that represents variables/functions in scope.
type Scope struct {
	Vals   map[string]interface{}
	Parent *Scope
	Files  map[string]*ast.File
	config *types.Config
	path   string
	line   int
	fset   *token.FileSet

	isSelect   bool
	typeAssert reflect.Type
	isFunction bool
	defers     []*Defer

	sync.Mutex
}

type Defer struct {
	fun       ast.Expr
	scope     *Scope
	arguments []interface{}
}

func (scope *Scope) Defer(d *Defer) error { _ = "STUB: not implemented"; return nil }

// NewScope creates a new initialized scope
func NewScope() *Scope { _ = "STUB: not implemented"; return nil }

// GetPointer walks the scope and finds the pointer to the value of interest
func (scope *Scope) GetPointer(name string) (val interface{}, exists bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Get walks the scope and finds the value of interest
func (scope *Scope) Get(name string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set walks the scope and sets a value in a parent scope if it exists, else current.
func (scope *Scope) Set(name string, val interface{}) { _ = "STUB: not implemented"; return }

// Keys returns all keys in scope
func (scope *Scope) Keys() (keys []string) { _ = "STUB: not implemented"; return nil }

// NewChild creates a scope under the existing scope.
func (scope *Scope) NewChild() *Scope { _ = "STUB: not implemented"; return nil }

// Func represents an interpreted function definition.
type Func struct {
	Def *ast.FuncLit
}

// ParseString parses go code into the ast nodes.
func (scope *Scope) ParseString(exprStr string) (ast.Node, int, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), 0, nil
}

// InterpretString interprets a string of go code and returns the result.
func (scope *Scope) InterpretString(exprStr string) (v interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Interpret interprets an ast.Node and returns the value.
func (scope *Scope) Interpret(expr ast.Node) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO make builtinScope root of other scopes

// Handle indirection cases.

// If not valid key, return the "zero" type. Eg for int 0, string ""

// TODO implement type checking
//define := e.Tok == token.DEFINE

// We're using a map here since we want iteration on clauses to be
// pseudo-random.

func (scope *Scope) getValue(id ast.Expr) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (scope *Scope) ExecuteFunc(funExpr ast.Expr, args []interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO enforce func return values

// ConfigureTypes configures the scope type checker
func (scope *Scope) ConfigureTypes(path string, line int) error {
	_ = "STUB: not implemented"
	return nil
}

// positions are relative to fset

// Parse the file containing this very example
// but stop after processing the imports.

// walker adapts a function to satisfy the ast.Visitor interface.
// The function return whether the walk should proceed into the node's children.
type walker func(ast.Node) bool

func (w walker) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

// CheckStatement checks if a statement is type safe
func (scope *Scope) CheckStatement(node ast.Node) (errs []error) {
	_ = "STUB: not implemented"
	return nil
}

// Render renders an ast node
func (scope *Scope) Render(x ast.Node) string { _ = "STUB: not implemented"; return "" }

// TypeCheck does type checking and returns the info object
func (scope *Scope) TypeCheck() (*types.Info, []error) { _ = "STUB: not implemented"; return nil, nil }

// these errors should be reported via the error reporter above

// StringToType returns the reflect.Type corresponding to the type string provided. Ex: StringToType("int")
func StringToType(str string) (reflect.Type, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), nil
}

// ValuesToInterfaces converts a slice of []reflect.Value to []interface{}
func ValuesToInterfaces(vals []reflect.Value) []interface{} { _ = "STUB: not implemented"; return nil }
