package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var addonsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an addon",
	Run: func(cmd *cobra.Command, args []string) {
		addonId, _ := cmd.Flags().GetInt("addon-id")
		registry, _ := cmd.Flags().GetString("registry")
		version, _ := cmd.Flags().GetString("version")
		pullPolicy, _ := cmd.Flags().GetString("pull-policy")
		activationPolicy, _ := cmd.Flags().GetString("activation-policy")

		addonSdk := addon.New(conf)
		payload := &addon.UpdateAddonInput{
			AddonId:          addonId,
			OciRegistry:      registry,
			OciVersion:       version,
			PullPolicy:       pullPolicy,
			ActivationPolicy: activationPolicy,
		}
		out, err := addonSdk.Update(payload)
		if err != nil {
			log.Fatal(err)
		}

		printer.Print(out)
	},
}

func init() {
	addonsUpdateCmd.Flags().Int("addon-id", 0, "Addon ID")
	_ = addonsUpdateCmd.MarkFlagRequired("addon-id")
	addonsUpdateCmd.Flags().String("registry", "", "OCI Url for the artifact")
	addonsUpdateCmd.Flags().String("version", "latest", "OCI version for the artifact")
	addonsUpdateCmd.Flags().String("pull-policy", "IfNotPresent", "Pull policy for the OCI artifact")
	addonsUpdateCmd.Flags().String("activation-policy", "Manual", "Crossplane package activation policy")

	addonsCmd.AddCommand(addonsUpdateCmd)
}
