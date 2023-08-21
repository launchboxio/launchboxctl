package main

import "github.com/spf13/cobra"

var projectsAttachAddonCmd = &cobra.Command{
	Use:   "attach-addon",
	Short: "Attach an addon to a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	projectsCmd.AddCommand(projectsAttachAddonCmd)
}
