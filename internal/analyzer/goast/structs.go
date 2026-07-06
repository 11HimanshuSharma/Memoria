package goast

import (
	"go/ast"
	"go/token"
	"strconv"
)

func (a *Analyzer) extractStructs(
	file *ast.File,
	result *Result,
) {

	for _, decl := range file.Decls {

		gen, ok := decl.(*ast.GenDecl)

		if !ok {
			continue
		}

		if gen.Tok != token.TYPE {
			continue
		}

		for _, spec := range gen.Specs {

			ts, ok := spec.(*ast.TypeSpec)

			if !ok {
				continue
			}

			st, ok := ts.Type.(*ast.StructType)

			if !ok {
				continue
			}

			item := Struct{

				Name: ts.Name.Name,

				Exported: ts.Name.IsExported(),
			}

			item.Fields =
				a.extractStructFields(st)

			result.Structs = append(
				result.Structs,
				item,
			)
		}
	}
}

func (a *Analyzer) extractStructFields(
	st *ast.StructType,
) []StructField {

	if st.Fields == nil {

		return nil
	}

	fields := make(
		[]StructField,
		0,
		len(st.Fields.List),
	)

	for _, field := range st.Fields.List {

		fieldType := nodeString(field.Type)

		tag := ""

		if field.Tag != nil {

			value, err := strconv.Unquote(
				field.Tag.Value,
			)

			if err == nil {

				tag = value
			}
		}

		// Embedded field
		if len(field.Names) == 0 {

			fields = append(
				fields,
				StructField{

					Name: fieldType,

					Type: fieldType,

					Embedded: true,

					Exported: ast.IsExported(fieldType),

					Tag: tag,
				},
			)

			continue
		}

		for _, name := range field.Names {

			fields = append(
				fields,
				StructField{

					Name: name.Name,

					Type: fieldType,

					Embedded: false,

					Exported: name.IsExported(),

					Tag: tag,
				},
			)
		}
	}

	return fields
}