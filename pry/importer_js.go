//go:build js
// +build js

package pry

import (
	"go/ast"
	"go/types"
)

func (s *Scope) parseDir() (map[string]*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getImporter() types.Importer { _ = "STUB: not implemented"; return *new(types.Importer) }

var defaultImporter = &JSImporter{
	packages: map[string]*types.Package{},
	Dir:      map[string]*ast.Package{},
}

func InternalSetImports(raw string) { _ = "STUB: not implemented"; return }
