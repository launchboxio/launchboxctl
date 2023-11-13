package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectInstallAddonCmd = &cobra.Command{
	Use:   "install-addon",
	Short: "Install an addon for a project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetInt("project-id")
		addonId, _ := cmd.Flags().GetInt("addon-id")

		input := &project.InstallAddonInput{
			ProjectId: projectId,
			AddonId:   addonId,
		}

		projectSdk := project.New(conf)
		out, err := projectSdk.InstallAddon(input)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(out)
	},
}

func init() {
	projectInstallAddonCmd.Flags().Int("project-id", 0, "Project ID")
	_ = projectInstallAddonCmd.MarkFlagRequired("project-id")

	projectInstallAddonCmd.Flags().Int("addon-id", 0, "Addon to install")
	_ = projectInstallAddonCmd.MarkFlagRequired("addon-id")

	projectsCmd.AddCommand(projectInstallAddonCmd)
}
