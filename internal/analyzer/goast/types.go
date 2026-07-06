package goast


type Result struct {
	PackageName string
	Imports []Import
	Functions []Function
	Structs []struct
	Interfaces []Interface
}

type Import struct {
	Path string
	Alias string
}

type Function struct {
	Name string
	Receiver string
	Exported bool
	Parameters []Parameter
	TypeParameters []TypeParameter
	Variadic bool
	Result []Parameter
}

type Struct struct {
	Declaration

	Fields []StructField
}

type StructField struct {
	Name string

	Type string

	Tag string

	Embedded bool

	Exported bool
}

type Interface struct {
	Declaration

	Methods []Function

	Embedded []string
}