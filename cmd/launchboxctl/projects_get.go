package main

import (
	"errors"
	"github.com/spf13/cobra"
)

var projectsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information for a single project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("please provide a project ID")
		}
		projectId := args[0]
		project, err := sdk.GetProject(projectId)
		if err != nil {
			return err
		}

		return printer.Output(project)
	},
}

func init() {
	projectsCmd.AddCommand(projectsGetCmd)
}
