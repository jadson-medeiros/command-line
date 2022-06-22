package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "Foo is a cool program",
	Long: `Foo is an extremely cool program.
         It will make your life easier and fun`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is where you implement the command
		fmt.Println("Cobra running...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
