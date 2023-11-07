package cmd

import "github.com/spf13/cobra"

var projectsCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Interact with clusters",
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
