package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login to LaunchboxHQ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running auth")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
