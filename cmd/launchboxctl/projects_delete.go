package main

import (
	"errors"
	"github.com/spf13/cobra"
)

var projectsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("please provide a project ID")
		}
		projectId := args[0]
		err := sdk.DeleteProject(projectId)
		return err
	},
}

func init() {
	projectsCmd.AddCommand(projectsDeleteCmd)
}
