package app

import (
	"github.com/11himanshusharma/memoria/internal/config"
	"github.com/11himanshusharma/memoria/internal/document"
	"github.com/11himanshusharma/memoria/internal/pipeline"
	"github.com/11himanshusharma/memoria/internal/repository"
	"github.com/11himanshusharma/memoria/internal/scanner"
)

type App struct {
	config     *config.Config
	repository *repository.Repository
	scanner    *scanner.Scanner
	loader     *document.Loader
	pipeline   *pipeline.Pipeline
}

func (a *App) Pipeline() *pipeline.Pipeline {
	return a.pipeline
}

func New(cfg *config.Config, repo *repository.Repository) *App {
	return &App{
		config:     cfg,
		repository: repo,
	}
}

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Repository() *repository.Repository {
	return a.repository
}