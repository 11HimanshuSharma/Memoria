package pipeline

import (
	"context"
)

type Pipeline struct {
	ctx *Context
}

func New(ctx *Context) *Pipeline {
	return &Pipeline{ctx: ctx}

}


func (p *Pipeline) Run() error {
	it, err := p.ctx.Scanner.Scan(p.ctx.Context, 
	p.ctx.Repository.Root)

	if err != nil {
		return err
	}
	for it.Next() {
		file := it.File()

		_, err := p.ctx.Loader.Load(
			p.ctx.Context,
			file,
		)
		if err != nil {
			return err
		}
	}
	return it.Err()
}