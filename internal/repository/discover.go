package repository

import (
	"fmt"
	"path/filepath"
)

func Discover(path string)(*Repository, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	repo := &Repository{}
	//Git discovery
	if root, ok, err := findGitRoot(abs); err != nil {
		return nil, err
	} else if ok {
		repo.Git = &GitInfo{
			Root: root,
		}
		repo.Root = root
		repo.Name = filepath.Base(root)
	}

	//go module discovery
	if root, moduleName, ok, err := findGoMod(abs); err != nil {
		return nil, err
	} else if ok {
		repo.Root = root
		repo.Name = filepath.Base(root)
		repo.Module = &Module{
			Name: moduleName,
		}
	}
	if repo.Root == "" {
		return nil, fmt.Errorf("repository not found")
	}
	return repo, nil
}