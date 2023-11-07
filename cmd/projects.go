package cmd

import "github.com/spf13/cobra"

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Interact with projects",
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
