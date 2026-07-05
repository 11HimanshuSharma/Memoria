package scanner

import (
	"context"
	"io/fs"
	"path/filepath"
)

func (s *Scanner) walk(ctx context.Context, root string) ([]SourceFile, error) {
	files := make([]SourceFile, 0, 512)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		
		if err != nil {
			return err
		}
		
		if d.IsDir() {
			if s.shouldSkipDirectory(d.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		
		info, err := d.Info()
		if err != nil {
			return err
		}
		
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		
		files = append(files, SourceFile{
			Path:         path,
			RelativePath: relative,
			Name:         d.Name(),
			Extension:    extension(path),
			Size:         info.Size(),
			ModifiedAt:   info.ModTime(),
			Mode:         info.Mode(),
		})
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return files, nil
}