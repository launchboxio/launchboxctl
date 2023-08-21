package main

import "github.com/spf13/cobra"

var addonsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available addons",
	RunE: func(cmd *cobra.Command, args []string) error {
		addons, err := sdk.ListAddons()
		if err != nil {
			return err
		}
		return printer.Output(addons)
	},
}

func init() {
	addonsCmd.AddCommand(addonsListCmd)
}
