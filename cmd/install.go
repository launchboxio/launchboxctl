package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Bootstrap a cluster to communicate with LaunchboxHQ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing launchbox...")

	},
}

func init() {

}
