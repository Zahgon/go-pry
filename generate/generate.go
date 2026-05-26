package generate

import (
	"context"
	"go/ast"

	"golang.org/x/tools/go/packages"
)

type Generator struct {
	contexts []pryContext
	debug    bool
	Config   packages.Config
}

func NewGenerator(debug bool) *Generator { _ = "STUB: not implemented"; return nil }

// Debug prints debug statements if debug is true.
func (g Generator) Debug(templ string, k ...interface{}) { _ = "STUB: not implemented"; return }

// ExecuteGoCmd runs the 'go' command with certain parameters.
func (g *Generator) ExecuteGoCmd(ctx context.Context, args []string, env []string) error {
	_ = "STUB: not implemented"
	return nil
}

// InjectPry walks the scope and replaces pry.Pry with pry.Apply(pry.Scope{...}).
func (g *Generator) InjectPry(filePath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// positions are relative to fset

// Parse the file containing this very example
// but stop after processing the imports.

// Print the imports from the file's AST.

// GetExports returns a string of gocode that represents the exports (constants/functions) of an ast.Package.
func (g *Generator) GetExports(importName string, files []*ast.File, added map[string]bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Print the imports from the file's AST.

/*
	case *ast.ValueSpec:
		if len(stmt.Values) > 0 {
			out, err := pry.InterpretExpr(scope, stmt.Values[0])
			if err != nil {
				fmt.Println("ERR", err)
				//continue
			} else {
				scope[obj.Name] = out
			}
		}
*/

//continue

// TODO Fix hack for very large constants

// GenerateFile generates a injected file.
func (g *Generator) GenerateFile(imports []string, extraStatements, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateAndExecuteFile generates and executes a temp file with the given imports
func (g *Generator) GenerateAndExecuteFile(ctx context.Context, imports []string, extraStatements string) error {
	_ = "STUB: not implemented"
	return nil
}

// RevertPry reverts the changes made by InjectPry.
func (g *Generator) RevertPry(modifiedFiles []string) error { _ = "STUB: not implemented"; return nil }

func filterVars(vars []string) (fVars []string) { _ = "STUB: not implemented"; return nil }

func (g *Generator) extractVariables(vars []string, l []ast.Stmt) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) extractFields(vars []string, l []*ast.Field) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleStatement(vars []string, s ast.Stmt) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleIfStmt(vars []string, stmt *ast.IfStmt) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleRangeStmt(vars []string, stmt *ast.RangeStmt) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleForStmt(vars []string, stmt *ast.ForStmt) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleBlockStmt(vars []string, stmt *ast.BlockStmt) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleIdents(vars []string, idents []*ast.Ident) []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *Generator) handleExpr(vars []string, v ast.Expr) []string {
	_ = "STUB: not implemented"
	return nil
}

//handleExpr(vars, fun.X)

type pryContext struct {
	Start, End int
	Vars       []string
}
