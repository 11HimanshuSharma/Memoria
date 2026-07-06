package goast


import (
	"go/ast"
	"strconv"
)


func (a *Analyzer) extractImports(file *ast.File, result *Result,) {
	if len(file.Imports) == 0 {
		return
	}
	result.Imports = make([]Import, 0, len(file.Imports))

	for _, imp := range file.Imports {
		path, err := strconv.Unquote(
			imp.Path.Value
		)
		if err != nil {
			continue

		}
		item := Import{
			Path: path,
		}
		if imp.Name != nil {
			item.Alias = imp.Name.Name
		}
		result.Imports = append(result.Imports,item)
	}
}