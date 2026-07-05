package cli


import ( 
	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use: "memoria",
	Short: "AI-native memory system for software engineering",
	Long: `Memoria continuosly transforms software development
	activity into a living engineering wiki`,
}


func Execute() error {
	return rootCmd.Execute()
}