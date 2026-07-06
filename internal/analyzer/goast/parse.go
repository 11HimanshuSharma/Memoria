package goast


import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"github.com/11himanshusharma/memoria/internal/document"
)

func (a *Analyzer) parseDocument(
	doc *document.Document,
	result *Result,
) error {
	fset := a.fset
	file, err := parser.ParseFile(
		fset,
		doc.File.Path,
		doc.Content,
		parser.ParseComments,
	)
	if err != nil {
		return fmt.Errorf("parse %s: %w", doc.File.RelativePath, err,)
	}
	return a.extractPasses(
		file, result,
	)
}

func (a *Analyzer) extractPasses(
	file *ast.File,
	result *Result,
) error {
	result.PackageName = file.Name.Name
	a.extractImports(
		file,
		result,
	)
	return nil
}