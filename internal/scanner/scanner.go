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
	ch := make(chan SourceFile, 128)
	
	go func() {
		_ = s.walk(ctx, abs, ch)
	} ()

	return newIterator(ctx, ch), nil
}
