package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

		out, err := json.MarshalIndent(project, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(out))
		return nil
	},
}

func init() {
	projectsCmd.AddCommand(projectsGetCmd)
}
