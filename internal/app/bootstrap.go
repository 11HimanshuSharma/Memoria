package app

import (
	"context"
	"fmt"

	"github.com/11himanshusharma/memoria/internal/config"
	"github.com/11himanshusharma/memoria/internal/document"
	"github.com/11himanshusharma/memoria/internal/pipeline"
	"github.com/11himanshusharma/memoria/internal/pipeline/stages"
	"github.com/11himanshusharma/memoria/internal/repository"
	"github.com/11himanshusharma/memoria/internal/scanner"
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
	
	sc := scanner.New()
	loader := document.NewLoader()
	pipe := pipeline.New(
		&pipeline.Runtime{
			Context:    context.Background(),
			Repository: repo,
			Scanner:    sc,
			Loader:     loader,
		},
	)
	pipe.Register(
		stages.NewScanner(),
	)

	return &App{
		config:     cfg,
		repository: repo,
		scanner:    sc,
		loader:     loader,
		pipeline:   pipe,
	}, nil
}
