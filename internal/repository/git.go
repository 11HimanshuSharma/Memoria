package repository

import (
	"os"
	"path/filepath"
)


//findGitRoot walks upwards until it finds a .git directory
func findGitRoot(start string) (string, bool, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", false, err
	}
	for {
		gitPath := filepath.Join(current, ".git")
		info, err := os.Stat(gitPath)

		if err == nil && info.IsDir() {
			return current, true, nil
		}
		parent := filepath.Dir(current)

		if parent == current {
			return "", false, nil
		}
		current = parent
	}
}