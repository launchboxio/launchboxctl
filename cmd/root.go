package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lbx",
	Short: "Command Line for LaunchboxHQ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launchbox is initialized, now with pre-commits!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
