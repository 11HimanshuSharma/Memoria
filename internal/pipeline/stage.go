package pipeline

import "context"

type Stage interface{
	Name() string
	Run(
		context.Context, 
		*Runtime,
	) error
}