package scanner
import "context"


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

func (s *Scanner) Scan(ctx context.Context, root string) ([]SourceFile, error) {
	return nil, nil
}
