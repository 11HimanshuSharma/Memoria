package goast


import (
	"go/ast"

)

func (a *Analyzer) extractFunctions(file *ast.File, result *Result,) {
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		function := Function{
			Name: fn.Name.Name
			Exported: fn.Name.IsExported(),
		}
		if fn.Recv != nil && len(fn.Recv.List) > 0 {
			
			receiver := parseField(fn.Recv.List[0],)
			
			function.Receiver = &Receiver{
				Name: receiver.Name,
				Type: receiver.Type,
			}
		}
		function.Parameters = parseFieldList(fn.Type.Params)

		function.Results = parseFieldList(fn.Type.Results)

		function.TypeParameters = parseTypeParameters(fn.Type.TypeParams)

		if len(fn.Type.Params.List) > 0 {
			last := fn.Type.Params.List[
				len(fn.Type.Params.List) - 1
			]
			_, function.Variadic = last.Type.(*ast.Ellipsis)
		}
		result.Functions = append(result.Functions, function)
	}
}