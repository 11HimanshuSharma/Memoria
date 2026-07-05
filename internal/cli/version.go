package cli

import (
	"fmt"
	"github.com/11himanshusharma/memoria/internal/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print Memoria version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version : %s\n", version.Version)
		fmt.Printf("Commit : %s\n", version.Commit)
		fmt.Printf("Date : %s\n", version.Date)
	},
}