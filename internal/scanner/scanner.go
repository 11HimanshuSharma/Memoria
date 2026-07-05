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
	it := newIterator(ctx)
	if err := s.walk(ctx, abs, it); err != nil {
		return nil, err
	}
	return it, nil
}
