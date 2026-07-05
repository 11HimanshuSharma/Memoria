package scanner

import (
	"os"
	"time"
)

type Entry struct {
	Path         string
	RelativePath string
	Info         os.FileInfo
	ModTime      time.Time
}
