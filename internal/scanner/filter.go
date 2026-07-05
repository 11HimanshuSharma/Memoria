package scanner

import (
	"strings"
)

var ignoreDirectories = map[string]struct{}{
	".git":         {},
	".memoria":     {},
	"node_modules": {},
	"vendor":       {},
	"dist":         {},
	"build":        {},
	"target":       {},
	".idea":        {},
	".vscode":      {},
}

func (s *Scanner) shouldSkipDirectory(name string) bool {
	if s.options.IgnoreGit {
		if _, ok := ignoreDirectories[name]; ok {
			return true
		}
	}
	if !s.options.IncludeHidden && strings.HasPrefix(name, ".") && name != "." {
		return true
	}
	return false
}

