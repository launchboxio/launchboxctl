package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Run: func(cmd *cobra.Command, args []string) {

		projectSdk := project.New(conf)
		projects, err := projectSdk.List(&project.ListProjectInput{})
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(projects)
	},
}
