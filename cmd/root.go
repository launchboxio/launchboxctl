package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/spf13/cobra"
	"os"
)

var (
	client *sling.Sling

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logo)

		cmd.Help()
	},
}

func init() {
	accessToken := os.Getenv("LBX_ACCESS_TOKEN")
	host := os.Getenv("LBX_ADDRESS")
	if host == "" {
		host = "https://api.launchboxhq.io"
	}
	client = sling.
		New().
		Base(fmt.Sprintf("%s/v1/", host)).
		Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func outputJson(in interface{}) error {
	b, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	fmt.Println(string(b))
	return nil
}
