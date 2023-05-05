package cmd

import "github.com/spf13/cobra"

var addonsCmd = &cobra.Command{
	Use:   "addons",
	Short: "Manage project scoped addons",
}

func init() {
	rootCmd.AddCommand(addonsCmd)
}
