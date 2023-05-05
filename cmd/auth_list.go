package cmd

import "github.com/spf13/cobra"

var authListCmd = &cobra.Command{
	Use:   "auth list",
	Short: "List configured connections",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
