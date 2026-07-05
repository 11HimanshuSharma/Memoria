package cli

import (
	"github.com/11himanshusharma/memoria/internal/app"
	"github.com/spf13/cobra"
)

var application *app.App

var rootCmd = &cobra.Command{
	Use:   "memoria",
	Short: "AI-native memory system for software engineering",
	Long: `Memoria continuosly transforms software development
	activity into a living engineering wiki`,
}

func Execute(a *app.App) error {
	application = a
	return rootCmd.Execute()
}
