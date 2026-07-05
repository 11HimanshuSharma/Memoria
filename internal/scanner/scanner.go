package scanner

import (
	"context"
	"path/filepath"
)


type Scanner struct {
	options Options
}

func New(opts ...Option) *Scanner {
	options := defaultOptions()

	for _, opt := range opts {
		opt(&options)
	}
	return &Scanner{
		options: options,
	}
}

func (s *Scanner) Scan(ctx context.Context, root string) (*Iterator, error) {
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	files := make(chan *SourceFile, 128)
	errors := make(chan error, 1)
	
	go func() {
		_ = s.walk(ctx, abs, files, errors)
	} ()

	return newIterator(ctx, files, errors), nil
}
