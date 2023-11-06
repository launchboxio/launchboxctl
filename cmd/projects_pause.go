package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectsPauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause a project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project-id")

		projectSdk := project.New(conf)
		projects, err := projectSdk.Pause(&project.PauseProjectInput{
			ProjectId: projectId,
		})
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(projects)
	},
}

func init() {
	projectsPauseCmd.Flags().Int("project-id", 0, "Project ID")

	projectsCmd.AddCommand(projectsPauseCmd)
}
