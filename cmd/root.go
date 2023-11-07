package cmd

import (
	"fmt"
	"github.com/launchboxio/launchbox-go-sdk/config"
	"github.com/spf13/cobra"
	"os"
)

var (
	conf *config.Config

	logo = `
 ___       ________  ___  ___  ________   ________  ___  ___  ________  ________     ___    ___
|\  \     |\   __  \|\  \|\  \|\   ___  \|\   ____\|\  \|\  \|\   __  \|\   __  \   |\  \  /  /|
\ \  \    \ \  \|\  \ \  \\\  \ \  \\ \  \ \  \___|\ \  \\\  \ \  \|\ /\ \  \|\  \  \ \  \/  / /
 \ \  \    \ \   __  \ \  \\\  \ \  \\ \  \ \  \    \ \   __  \ \   __  \ \  \\\  \  \ \    / /
  \ \  \____\ \  \ \  \ \  \\\  \ \  \\ \  \ \  \____\ \  \ \  \ \  \|\  \ \  \\\  \  /     \/
   \ \_______\ \__\ \__\ \_______\ \__\\ \__\ \_______\ \__\ \__\ \_______\ \_______\/  /\   \
    \|_______|\|__|\|__|\|_______|\|__| \|__|\|_______|\|__|\|__|\|_______|\|_______/__/ /\ __\
                                                                                    |__|/ \|__|
`
)

var rootCmd = &cobra.Command{
	Use:   "lbx",
	Short: "Command Line for LaunchboxHQ",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		endpoint, _ := cmd.PersistentFlags().GetString("endpoint")
		var err error
		if endpoint == "" {
			conf, err = config.Default()
		} else {
			conf, err = config.DefaultWithEndpoint(endpoint)
		}
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logo)

		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().String("endpoint", "", "API Endpoint for LaunchboxHQ")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
