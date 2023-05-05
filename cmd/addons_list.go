package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var addonsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List addons",
	Run: func(cmd *cobra.Command, args []string) {
		addons := new([]Addon)
		_, err := client.Get("addons").ReceiveSuccess(addons)
		if err != nil {
			log.Fatal(err)
		}
		_ = outputJson(addons)
	},
}
