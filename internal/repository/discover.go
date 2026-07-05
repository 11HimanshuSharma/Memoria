package repository

import (
	"fmt"
	"path/filepath"
)

func Discover(path string) (*Context, error) {
	ctx := &Context{}

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	root, ok, err := findGitRoot(abs)
	if err != nil {
		return nil, err
	}

	if ok {
		ctx.IsGit = true
		ctx.Root = root
		ctx.Name = filepath.Base(root)
	}

	// -------------
	goRoot, module, ok, err := findGoMod(abs)
	if err != nil {
		return nil, err
	}
	if ok {
		ctx.Root = goRoot
		ctx.Name = filepath.Base(goRoot)

		ctx.Module = module
		ctx.Language = Go
		ctx.BuildSystem = GoModules
	}

	if ctx.Root == "" {
		return nil, fmt.Errorf("repository not found")
	}
	return ctx, nil
}