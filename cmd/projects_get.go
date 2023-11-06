package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project-id")

		projectSdk := project.New(conf)
		project, err := projectSdk.Get(&project.GetProjectInput{
			ProjectId: projectId,
		})
		if err != nil {
			log.Fatal(err)
		}
		err = printer.Print(project)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	projectsGetCmd.Flags().Int("project-id", 0, "Project ID")

	projectsCmd.AddCommand(projectsGetCmd)
}
