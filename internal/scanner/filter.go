package scanner

func (s *Scanner) shouldSkipDirectory(name string) bool {

	if s.policy.IgnoreDirectories &&
		s.ignore.IgnoreDir(name) {

		return true
	}

	if !s.policy.IncludeHidden &&
		s.ignore.Hidden(name) {

		return true
	}

	return false
}

func (s *Scanner) shouldSkipFile(name string) bool {

	if s.policy.IgnoreFiles &&
		s.ignore.IgnoreFile(name) {

		return true
	}

	if !s.policy.IncludeHidden &&
		s.ignore.Hidden(name) {

		return true
	}

	return false
}