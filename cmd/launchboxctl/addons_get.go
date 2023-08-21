package main

import (
	"errors"
	"github.com/spf13/cobra"
)

var addonsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve data about an addon",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("please provide an Addon ID")
		}
		addonId := args[0]
		addon, err := sdk.GetAddon(addonId)
		if err != nil {
			return err
		}
		return printer.Output(addon)
	},
}

func init() {
	addonsCmd.AddCommand(addonsGetCmd)
}
