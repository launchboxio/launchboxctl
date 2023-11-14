package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectUninstallAddonCmd = &cobra.Command{
	Use:   "uninstall-addon",
	Short: "Uninstall an addon for a project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project-id")
		subscriptionId, _ := cmd.Flags().GetInt("subscription-id")

		input := &project.UninstallAddonInput{
			ProjectId:      projectId,
			SubscriptionId: subscriptionId,
		}

		projectSdk := project.New(conf)
		out, err := projectSdk.UninstallAddon(input)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(out)
	},
}

func init() {
	projectUninstallAddonCmd.Flags().Int("project-id", 0, "Project ID")
	_ = projectUninstallAddonCmd.MarkFlagRequired("project-id")

	projectUninstallAddonCmd.Flags().Int("subscription-id", 0, "ID of the addon subscription")
	_ = projectUninstallAddonCmd.MarkFlagRequired("subscription-id")

	projectsCmd.AddCommand(projectUninstallAddonCmd)
}
