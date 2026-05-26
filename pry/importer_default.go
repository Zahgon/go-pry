//go:build !js
// +build !js

package pry

import (
	"go/ast"
	"go/types"
)

type packagesImporter struct {
}

func (i packagesImporter) Import(path string) (*types.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (packagesImporter) ImportFrom(path, dir string, mode types.ImportMode) (*types.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getImporter() types.ImporterFrom { _ = "STUB: not implemented"; return *new(types.ImporterFrom) }

func (s *Scope) parseDir() (map[string]*ast.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
