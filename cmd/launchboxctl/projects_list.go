package main

import (
	"encoding/json"
	"fmt"
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

		out, err := json.MarshalIndent(projects, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(out))
		return nil
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
