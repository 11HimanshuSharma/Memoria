package repository

type Context struct {
	Name        string
	Root        string
	Module      string
	Language    Language
	BuildSystem BuildSystem
	IsGit       bool
}
