package stages
import (
	"context"
	"github.com/11himanshusharma/memoria/internal/pipeline"
)

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) Name() string {
	return "scanner"
}

func (s *Scanner) Run(
	ctx context.Context,
	rt *pipeline.Runtime,
) error {
	_, err := rt.Scanner.Scan(
		ctx, 
		rt.Repository.Root,
	)
	return err
}
