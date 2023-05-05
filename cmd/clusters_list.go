package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var clustersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List clusters",
	Run: func(cmd *cobra.Command, args []string) {
		clusters := new([]Cluster)
		_, err := client.Get("clusters").ReceiveSuccess(clusters)
		if err != nil {
			log.Fatal(err)
		}
		_ = outputJson(clusters)
	},
}
