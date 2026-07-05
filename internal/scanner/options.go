package scanner

type Options struct {
	Workers         int
	FollowSymlinks  bool
	IncludeHidden   bool
	IgnoreGit       bool
	IgnoreGitIgnore bool
}

func defaultOptions() Options {
	return Options{
		Workers: 8,
		FollowSymlinks: false,
		IncludeHidden: false,
		IgnoreGit: true,
		IgnoreGitIgnore: true,
	}
}

type Option func(*Options)

func WithWorkers(n int) Option {
	return func(o *Options) {
		if n > 0 {
			o.Workers = n
		}
	}
}

func WithFollowSymlinks(enable bool) Option {
	return func(o *Options) {
		o.FollowSymlinks  = enable
	}
}
func WithHiddenFiles(enable bool) Option {
	return func(o *Options) {
		o.IncludeHidden = enable
	}
}
func WithGitDirectory(enable bool) Option {
	return func(o *Options) {
		o.IgnoreGit = !enable
	}
}
func WithGitIgnore(enable bool) Option {
	return func(o *Options) {
		o.IgnoreGitIgnore = enable
	}
}