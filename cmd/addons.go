package cmd

import "github.com/spf13/cobra"

type Addon struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	OciRegistry string `json:"oci_registry,omitempty"`
	OciVersion  string `json:"oci_version,omitempty"`
}

var addonsCmd = &cobra.Command{
	Use:   "addons",
	Short: "Manage project scoped addons",
}

func init() {
	rootCmd.AddCommand(addonsCmd)

	addonsCmd.AddCommand(addonsListCmd)
	addonsCmd.AddCommand(addonsCreateCmd)
	addonsCmd.AddCommand(addonsGetCmd)
}
