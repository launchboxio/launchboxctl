package cmd

import "github.com/spf13/cobra"

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage LaunchboxHQ connections",
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(authAddCmd)
	authCmd.AddCommand(authListCmd)
}
