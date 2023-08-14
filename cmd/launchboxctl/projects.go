package main

import "github.com/spf13/cobra"

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage Launchbox projects",
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
