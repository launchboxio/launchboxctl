package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectListAddonsCmd = &cobra.Command{
	Use:   "list-addons",
	Short: "List currently installed addons",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project-id")

		input := &project.ListAddonsInput{
			ProjectId: projectId,
		}

		projectSdk := project.New(conf)
		out, err := projectSdk.ListAddons(input)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(out)
	},
}

func init() {
	projectListAddonsCmd.Flags().Int("project-id", 0, "Project ID")
	_ = projectListAddonsCmd.MarkFlagRequired("project-id")

	projectsCmd.AddCommand(projectListAddonsCmd)
}
