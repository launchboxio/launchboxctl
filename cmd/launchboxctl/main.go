package main

import (
	"github.com/launchboxio/launchboxctl/pkg/client"
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:          "lbx",
		Short:        "LaunchboxHQ CLI",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			sdk = client.New(&client.ClientOpts{
				Host:  host,
				Token: token,
			})
		},
	}

	host     string
	clientId string
	token    string
	sdk      *client.Client
)

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "https://launchboxhq.io", "Launchbox Installation Host")
	rootCmd.PersistentFlags().StringVar(&clientId, "client-id", "", "OIDC Client ID for authentication")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
