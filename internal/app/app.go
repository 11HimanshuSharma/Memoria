package app

import (
	"github.com/11himanshusharma/memoria/internal/config"
	"github.com/11himanshusharma/memoria/internal/repository"
)

type App struct {
	config *config.Config
	repository *repository.Context
}

func New(cfg *config.Config, repo *repository.Context) *App {
	return &App{
		config: cfg,
		repository: repo,
	}

}

func (a *App) Config() *config.Config {
	return a.config
}


func (a *App) Repository() *repository.Context {
	return a.repository
}