package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var addonsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an addon",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Please pass the ID as the first argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		_, err := client.
			Delete(fmt.Sprintf("addons/%s", id)).
			ReceiveSuccess(nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Addon deleted")
	},
}
