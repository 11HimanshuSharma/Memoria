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
		fmt.Println(" Go Version: ", runtime.Version())
		fmt.Println(" OS: ", runtime.GOOS)
		fmt.Println(" Architecture", runtime.GOARCH)
	},
}
