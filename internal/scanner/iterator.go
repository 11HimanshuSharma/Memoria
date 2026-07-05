package scanner

import "context"

type Iterator struct {
	ctx context.Context
	files <- chan *SourceFile
	current *SourceFile
	err error
}

func newIterator(
	ctx context.Context,
	files <-chan *SourceFile,
	errors <- chan error,
) *Iterator {
	return &Iterator{
		ctx: ctx,
		files: files,
	}
}

func (it *Iterator) Next() bool {
	select {
	case <- it.ctx.Done():
		it.err = it.ctx.Err()
		return false
	case file, ok := <-it.files:
		if !ok {
			return false
		}
		it.current = file
		return true
	}
}

func (it *Iterator) File() *SourceFile {
	return it.current
}
func (it *Iterator) Err() error {
	return it.err
}