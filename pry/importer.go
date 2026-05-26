package pry

import (
	"go/ast"
	"go/types"
)

// JSImporter contains all the information needed to implement a types.Importer
// in a javascript environment.
type JSImporter struct {
	packages map[string]*types.Package
	Dir      map[string]*ast.Package
}

func (i *JSImporter) Import(path string) (*types.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
