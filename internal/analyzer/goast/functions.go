package goast

import (
	"go/ast"
)

func (a *Analyzer) extractFunctions(
	file *ast.File,
	result *Result,
) {

	for _, decl := range file.Decls {

		fn, ok := decl.(*ast.FuncDecl)

		if !ok {
			continue
		}

		function := Function{
			Name:      fn.Name.Name,
			Exported:  fn.Name.IsExported(),
			Parameters: parseFieldList(fn.Type.Params),
			Results:    parseFieldList(fn.Type.Results),
			TypeParameters: parseTypeParameters(fn.Type.TypeParams),
		}

		// Receiver (method)
		if fn.Recv != nil && len(fn.Recv.List) > 0 {

			recv := parseField(fn.Recv.List[0])

			function.Receiver = &Receiver{
				Name: recv.Name,
				Type: recv.Type,
			}
		}

		// Variadic
		if fn.Type.Params != nil &&
			len(fn.Type.Params.List) > 0 {

			last := fn.Type.Params.List[len(fn.Type.Params.List)-1]

			if _, ok := last.Type.(*ast.Ellipsis); ok {
				function.Variadic = true
			}
		}

		result.Functions = append(
			result.Functions,
			function,
		)
	}
}