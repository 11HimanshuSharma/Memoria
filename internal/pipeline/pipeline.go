package pipeline

type Pipeline struct {
	runtime *Runtime
	stages  []Stage
}

func New(runtime *Runtime) *Pipeline {
	return &Pipeline{
		runtime: runtime,
		stages:  make([]Stage, 0),
	}

}
func (p *Pipeline) Register(stage Stage) {
	p.stages = append(p.stages, stage)
}

func (p *Pipeline) Run() error {
	for _, stage := range p.stages {
		err := stage.Run(
			p.runtime.Context,
			p.runtime,
		)
		if err != nil {
			return err
		}
	}
	return nil
}