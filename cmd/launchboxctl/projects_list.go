package main

import (
	"github.com/spf13/cobra"
)

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List LaunchboxHQ projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		projects, err := sdk.ListProjects()
		if err != nil {
			return err
		}

		return printer.Output(projects)
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
