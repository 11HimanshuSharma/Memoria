package pipeline

import (
	"context"
	"github.com/11himanshusharma/memoria/internal/document"
	"github.com/11himanshusharma/memoria/internal/repository"
	"github.com/11himanshusharma/memoria/internal/scanner"
)

type Runtime struct {
	Context context.Context
	Repository *repository.Repository
	Scanner *scanner.Scanner
	Loader *document.Loader
}
