package repository

type Repository struct {
	Name string
	Root string
	Module *Module
	Git *GitInfo
}