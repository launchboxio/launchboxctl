package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectsResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Resume a project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project-id")

		projectSdk := project.New(conf)
		projects, err := projectSdk.Resume(&project.ResumeProjectInput{
			ProjectId: projectId,
		})
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(projects)
	},
}

func init() {
	projectsResumeCmd.Flags().Int("project-id", 0, "Project ID")

	projectsCmd.AddCommand(projectsResumeCmd)
}
