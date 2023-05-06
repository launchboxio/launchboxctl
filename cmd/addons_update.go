package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var addonsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an addon",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Please pass the ID as the first argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		addon := &Addon{}

		name, _ := cmd.Flags().GetString("name")
		registry, _ := cmd.Flags().GetString("registry")
		version, _ := cmd.Flags().GetString("version")

		if name != "" {
			addon.Name = name
		}
		if registry != "" {
			addon.OciRegistry = registry
		}
		if version != "" {
			addon.OciVersion = version
		}

		if addon == (&Addon{}) {
			fmt.Println("No changes found")
			return
		}
		_, err := client.
			Put(fmt.Sprintf("addons/%s", id)).
			BodyJSON(addon).
			ReceiveSuccess(addon)
		if err != nil {
			log.Fatal(err)
		}

		_ = outputJson(addon)
	},
}

func init() {
	addonsUpdateCmd.Flags().String("registry", "", "OCI Url for the artifact")
	addonsUpdateCmd.Flags().String("version", "", "OCI Version for the artifact")
	addonsUpdateCmd.Flags().String("name", "", "Name of the addon")
}
