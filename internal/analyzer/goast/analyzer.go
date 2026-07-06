package goast

import (
	"github.com/11himanshusharma/memoria/internal/document"
)

type Analyzer struct {}

func New() *Analyzer{
	return &Analyzer{}
}

func (a *Analyzer) Analyze(doc *document.Document,) (*Result, error) {
	result := &Result{}
	err := a.parseDocument(
		doc, result,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
