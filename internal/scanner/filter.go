package scanner

import (
	"strings"
)


func (s *Scanner) shouldSkipDirectory(name string) bool {
	if s.options.IgnoreGit && s.ignore.IgnoreDir(name) {
		return true
	}
	if !s.options.IncludeHidden && s.ignore.Hidden(name) {
		return true
	}
	return false
}


func (s *Scanner) shouldSkipFile(name string) bool {
	if s.ignore.IgnoreFile(name) {
		return true
	}
	if !s.options.IncludeHidden && s.ignore.Hidden(name) {
		return true
	}
	return false

}
