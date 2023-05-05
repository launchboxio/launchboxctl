package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var authAddCmd = &cobra.Command{
	Use:   "auth add",
	Short: "Add a new connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter your launchbox domain [api.launchboxhq.io]:")
		var domain string
		fmt.Scanln(&domain)

		if domain == "" {
			domain = "api.launchboxhq.io"
		}

		fmt.Println("Opening a browser for authentication..")

		fmt.Println("Enter your access token:")
		var accessToken string
		fmt.Scanln(&accessToken)

		fmt.Printf("Access token %s for host %s", accessToken, domain)
	},
}
