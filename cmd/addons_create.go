package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"github.com/launchboxio/launchboxctl/internal/printer"
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
		pullPolicy, _ := cmd.Flags().GetString("pull-policy")
		activationPolicy, _ := cmd.Flags().GetString("activation-policy")

		input := &addon.CreateAddonInput{
			Name:             name,
			OciRegistry:      registry,
			OciVersion:       version,
			PullPolicy:       pullPolicy,
			ActivationPolicy: activationPolicy,
		}

		addonSdk := addon.New(conf)
		out, err := addonSdk.Create(input)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(out)
	},
}

func init() {
	addonsCreateCmd.Flags().String("name", "", "Addon name")
	_ = addonsCreateCmd.MarkFlagRequired("name")

	addonsCreateCmd.Flags().String("registry", "", "OCI Url for the artifact")
	_ = addonsCreateCmd.MarkFlagRequired("registry")

	addonsCreateCmd.Flags().String("version", "latest", "OCI version for the artifact")
	addonsCreateCmd.Flags().String("pull-policy", "IfNotPresent", "Pull policy for the OCI artifact")
	addonsCreateCmd.Flags().String("activation-policy", "Manual", "Crossplane package activation policy")

	addonsCmd.AddCommand(addonsCreateCmd)
}
