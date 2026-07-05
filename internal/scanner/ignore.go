package scanner

import (
	"path/filepath"
	"strings"
)

type IgnoreEngine struct {
	directories map[string]struct{}
	files  		map[string]struct{}
}

func NewIgnoreEngine() *IgnoreEngine {
	return &IgnoreEngine{
		directories: map[string]struct{}{
			".git":         {},
			".memoria":     {},
			"node_modules": {},
			"vendor":       {},
			"dist":         {},
			"build":        {},
			"target":       {},
			".idea":        {},
			".vscode":      {},
		},
		files: map[string]struct{}{
			".DS_Store": {},
		},
	}
}

func (e *IgnoreEngine) IgnoreDir(name string) bool {
	_, ok := e.directories[name]
	return ok
}
func (e *IgnoreEngine) IgnoreFile(name string) bool {
	_, ok := e.files[name]
	return ok
}

func (e *IgnoreEngine) Hidden(name string) bool {
	return strings.HasPrefix(name, ".") && name != "."
}
func (e *IgnoreEngine) Extension(path string) string {
	return strings.ToLower(filepath.Ext(path))
}
