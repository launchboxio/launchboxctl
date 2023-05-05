package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects := new([]Project)
		_, err := client.Get("projects").ReceiveSuccess(projects)
		if err != nil {
			log.Fatal(err)
		}
		_ = outputJson(projects)
	},
}
