package app

import (
	"fmt"
	"github.com/11himanshusharma/memoria/internal/config"
	"github.com/11himanshusharma/memoria/internal/repository"
)

func Bootstrap(configPath string) (*App, error) {
	cfg, err := config.Load(configPath)

	if err != nil {
		return nil, fmt.Errorf("Load config: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}
	repo, err := repository.Discover(cfg.Project.Path)
	if err != nil {
		return nil, fmt.Errorf("discover repository: %w", err)
	}

	return New(cfg, repo), nil
}
