package main

import "github.com/spf13/cobra"

var projectsAddonsCmd = &cobra.Command{
	Use:   "get-addons",
	Short: "List addons attached to a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		projectId, _ := cmd.Flags().GetInt("project-id")
		addons, err := sdk.GetProjectAddons(projectId)
		if err != nil {
			return err
		}
		return printer.Output(addons)
	},
}

func init() {
	projectsAddonsCmd.Flags().Int("project-id", 0, "ID of the project")
	projectsAddonsCmd.MarkFlagRequired("project-id")

	projectsCmd.AddCommand(projectsAddonsCmd)
}
