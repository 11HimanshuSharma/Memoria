package scanner

import "errors"

var (
	ErrOutsideRepository = errors.New("path is outside repository")
	ErrSymlinkLoop = errors.New("symlink look detected")
	ErrPermissionDenied = errors.New("permission denied")
)
