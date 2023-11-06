package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var addonsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List addons",
	Run: func(cmd *cobra.Command, args []string) {
		addonSdk := addon.New(conf)
		addons, err := addonSdk.List(&addon.ListAddonInput{})
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(addons)
	},
}
