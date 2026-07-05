package scanner

import (
	"os"
	"path/filepath"
	"time"
)

type SourceFile struct {
	// absolute path on disk
	Path string

	//path relative to repository root
	RelativePath string

	// file name
	Name string

	//Extension (.go, .md)
	Extension string

	//Size in bytes
	Size int64


	// Last  modified time
	ModifiedAt time.Time

	//Cached file mode
	Mode os.FileMode
}

func (f SourceFile) IsGo() bool {
	return f.Extension == ".go"
}
func (f SourceFile) IsMarkDown() bool {
	return f.Extension == ".md"
}

func (f SourceFile) Dir() string {
	return filepath.Dir(f.RelativePath)
}