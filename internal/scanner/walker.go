package scanner

import (
	"context"
	"io/fs"
	"path/filepath"
)

func (s *Scanner) walk(ctx context.Context, root string, out chan<- *SourceFile, errs chan<- error) error {
	defer close(out)
	defer close(errs)

	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
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
		if s.shouldSkipDirectory(d.Name()) {
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

		entry := Entry{
			Path:         path,
			RelativePath: relative,
			Info:         info,
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- s.process(entry):
		}
		return nil
	})
}
