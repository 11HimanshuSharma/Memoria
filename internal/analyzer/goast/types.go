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
}

type Struct struct {
	Name string
	Exported bool
}

type Interface struct {
	Name string
	Exported bool
}