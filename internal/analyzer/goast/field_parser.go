package goast

import (
	"bytes"
	"go/ast"
	"go/format"
)

func parseField(field *ast.Field) Parameter {

	param := Parameter{}

	if field == nil {
		return param
	}

	if len(field.Names) > 0 {
		param.Name = field.Names[0].Name
	}

	param.Type = nodeString(field.Type)

	return param
}

func parseFieldList(list *ast.FieldList) []Parameter {

	if list == nil {
		return nil
	}

	parameters := make([]Parameter, 0, len(list.List))

	for _, field := range list.List {

		parameter := parseField(field)

		// unnamed parameter
		if len(field.Names) == 0 {

			parameters = append(
				parameters,
				parameter,
			)

			continue
		}

		// one AST field may declare multiple names
		for _, name := range field.Names {

			item := parameter
			item.Name = name.Name

			parameters = append(
				parameters,
				item,
			)
		}
	}

	return parameters
}

func parseTypeParameters(
	list *ast.FieldList,
) []TypeParameter {

	if list == nil {
		return nil
	}

	parameters := make(
		[]TypeParameter,
		0,
		len(list.List),
	)

	for _, field := range list.List {

		constraint := nodeString(field.Type)

		for _, name := range field.Names {

			parameters = append(
				parameters,
				TypeParameter{
					Name: name.Name,

					Constraint: constraint,
				},
			)
		}
	}

	return parameters
}

func nodeString(node ast.Node) string {

	if node == nil {
		return ""
	}

	var buffer bytes.Buffer

	if err := format.Node(
		&buffer,
		nil,
		node,
	); err != nil {

		return ""
	}

	return buffer.String()
}