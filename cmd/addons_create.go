package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var addonsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an addon",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		registry, _ := cmd.Flags().GetString("registry")
		version, _ := cmd.Flags().GetString("version")

		addon := &Addon{
			Name:        name,
			OciRegistry: registry,
			OciVersion:  version,
		}

		_, err := client.Post("addons").BodyJSON(addon).ReceiveSuccess(addon)
		if err != nil {
			log.Fatal(err)
		}

		_ = outputJson(addon)
	},
}

func init() {
	addonsCreateCmd.Flags().String("name", "", "Addon name")
	_ = addonsCreateCmd.MarkFlagRequired("name")

	addonsCreateCmd.Flags().String("registry", "", "OCI Url for the artifact")
	_ = addonsCreateCmd.MarkFlagRequired("registry")

	addonsCreateCmd.Flags().String("version", "latest", "OCI version for the artifact")
}
