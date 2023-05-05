package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var addonsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "View an addon",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Please pass the ID as the first argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		addon := &Addon{}
		id := args[0]
		_, err := client.Get(fmt.Sprintf("addons/%s", id)).ReceiveSuccess(addon)
		if err != nil {
			log.Fatal(err)
		}

		_ = outputJson(addon)
	},
}
