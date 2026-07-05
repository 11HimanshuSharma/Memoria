package app

import (
	"github.com/11himanshusharma/memoria/internal/config"
	"github.com/11himanshusharma/memoria/internal/repository"
)

type App struct {
	config *config.Config
	repository *repository.Repository
}

func New(cfg *config.Config, repo *repository.Repository) *App {
	return &App{
		config: cfg,
		repository: repo,
	}

}

func (a *App) Config() *config.Config {
	return a.config
}


func (a *App) Repository() *repository.Repository {
	return a.repository
}