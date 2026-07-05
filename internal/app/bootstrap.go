package app

import (
	"fmt"
	"github.com/11himanshusharma/memoria/internal/config"
)

func Bootstrap(configPath string) (*App, error) {
	cfg, err := config.Load(configPath)

	if err != nil {
		return nil, fmt.Errorf("Load config: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}
	
	return &App{
		Config: cfg,

	}, nil
}