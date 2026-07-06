package scanner

type Policy struct {
	Workers int
	FollowSymlinks bool
	IncludeHidden bool
	IgnoreDirectories bool
	IgnoreFiles bool
	UseGitIgnore bool
}


func DefaultPolicy() Policy {
	return Policy{
		Workers: 8,
		FollowSymlinks: false,
		IncludeHidden: false,
		IgnoreDirectories: true,
		IgnoreFiles: true,
		UseGitIgnore: false,
	}
}

type Option func(*Policy)

func WithWorkers(n int) Option {
	return func(p *Policy) {
		if n > 0 {
			p.Workers = n
		}
	}
}
func WithHiddenFiles(enable bool) Option {
	return func(p *Policy) {
		p.IncludeHidden = enable
	}
}

func WithDirectoryFiltering(enable bool) Option {
	return func(p *Policy) {
		p.IgnoreDirectories = enable
	}
}

func WithFileFiltering(enable bool) Option {
	return func(p *Policy) {
		p.IgnoreFiles = enable
	}
}
func WithGitIgnore(enable bool) Option {
	return func(p *Policy) {
		p.UseGitIgnore = enable
	}
}