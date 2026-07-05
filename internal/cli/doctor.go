package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

func init() {
	rootCmd.AddCommand(doctorCmd)
}

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check Memoria installation",
	Run: func(cmd *cobra.Command, args []string) {

		repo := application.Repository()

		fmt.Println("Memoria Doctor")
		fmt.Println("──────────────────────────────")

		fmt.Println()
		fmt.Println("Runtime")
		fmt.Println("──────────────────────────────")
		fmt.Println("Go Version :", runtime.Version())
		fmt.Println("OS         :", runtime.GOOS)
		fmt.Println("Arch       :", runtime.GOARCH)


		fmt.Println()

		fmt.Println("Repository")
		fmt.Println("──────────────────────────────")
		fmt.Println("Name         :", repo.Name)
		fmt.Println("Root         :", repo.Root)
		fmt.Println("Module       :", repo.Module)
		fmt.Println("Language     :", repo.Language)
		fmt.Println("Build System :", repo.BuildSystem)
		fmt.Println("Git          :", repo.IsGit)
	},
}
