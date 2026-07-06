package document

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/11himanshusharma/memoria/internal/scanner"

)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}


func (l *Loader) Load(ctx context.Context, file *scanner.SourceFile) (*Document, error) {
	select {
	case <- ctx.Done():
		return nil, ctx.Err()
	default:
	}
	content, err := os.ReadFile(file.Path)
	if err != nil {
		return nil, fmt.Errorf(
			"read %s: %w",
			file.Path,
			err,
		)
	}
	return &Document{
		File: file,
		Content: content,
		LoadedAt: time.Now(),
	}, nil
}