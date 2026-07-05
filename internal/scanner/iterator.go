package scanner

import "context"

type Iterator struct {
	ctx   context.Context
	files []SourceFile
	index int
	err   error
}

func newIterator(ctx context.Context) *Iterator {
	return &Iterator{
		ctx:   ctx,
		files: make([]SourceFile, 0, 512),
		index: -1,
	}
}

func (it *Iterator) add(file SourceFile) {
	it.files = append(it.files, file)
}

func (it *Iterator) Next() bool {
	select {
	case <-it.ctx.Done():
		it.err = it.ctx.Err()
		return false
	default:
	}
	it.index++
	return it.index < len(it.files)
}

func (it *Iterator) File() SourceFile {
	return it.files[it.index]
}

func (it *Iterator) Err() error {
	return it.err
}