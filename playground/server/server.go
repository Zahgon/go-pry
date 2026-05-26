package main

import (
	"flag"
	"log"
	"net/http"
)

const bundlesDir = "bundles"

var bind = flag.String("bind", ":8080", "address to bind to")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func pkgHash(pkgs []string) string { _ = "STUB: not implemented"; return "" }

func normalizePackages(packages string) []string { _ = "STUB: not implemented"; return nil }

func generateBundle(w http.ResponseWriter, r *http.Request, packages string) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

func run() error { _ = "STUB: not implemented"; return nil }
