package document

import (
	"time"
	"github.com/11himanshusharma/memoria/internal/scanner"
)

type Document struct {
	// original scanned file
	File *scanner.SourceFile

	// Entire file contents
	Content []byte

	//Cached metadata
	LoadedAt time.Time
}

func (d *Document) String() string{
	return d.File.RelativePath
}