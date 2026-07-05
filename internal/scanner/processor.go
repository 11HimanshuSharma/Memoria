package scanner

import "path/filepath"

func (s *Scanner) process(entry Entry) *SourceFile {
	info := entry.Info

	return &SourceFile{
		Path:         entry.Path,
		RelativePath: entry.RelativePath,
		Name:         info.Name(),
		Extension:    filepath.Ext(entry.Path),
		Size:         info.Size(),
		ModifiedAt:   info.ModTime(),
		Mode:         info.Mode(),
	}
}