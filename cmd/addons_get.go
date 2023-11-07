package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var addonsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "View an addon",
	Run: func(cmd *cobra.Command, args []string) {
		addonId, _ := cmd.Flags().GetInt("addon-id")

		addonSdk := addon.New(conf)
		out, err := addonSdk.Get(&addon.GetAddonInput{
			AddonId: addonId,
		})
		if err != nil {
			log.Fatal(err)
		}

		printer.Print(out)
	},
}

func init() {
	addonsGetCmd.Flags().Int("addon-id", 0, "ID of the Addon")
	addonsCmd.AddCommand(addonsGetCmd)
}
