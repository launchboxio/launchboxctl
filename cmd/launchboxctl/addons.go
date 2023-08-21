package main

import "github.com/spf13/cobra"

var addonsCmd = &cobra.Command{
	Use:   "addons",
	Short: "Manage launchbox addons",
}

func init() {
	rootCmd.AddCommand(addonsCmd)
}
