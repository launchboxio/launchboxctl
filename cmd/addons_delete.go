package cmd

import (
	"fmt"
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"github.com/spf13/cobra"
	"log"
)

var addonsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an addon",
	Run: func(cmd *cobra.Command, args []string) {
		addonId, _ := cmd.Flags().GetInt("addon-id")

		addonSdk := addon.New(conf)
		_, err := addonSdk.Delete(&addon.DeleteAddonInput{
			AddonId: addonId,
		})
		if err != nil {
			log.Fatal(nil)
		}
		fmt.Println("Addon deleted")
	},
}

func init() {
	addonsDeleteCmd.Flags().Int("addon-id", 0, "Addon ID")
	_ = addonsDeleteCmd.MarkFlagRequired("addon-id")

	addonsCmd.AddCommand(addonsDeleteCmd)
}
