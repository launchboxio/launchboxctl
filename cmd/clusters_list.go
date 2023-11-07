package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/cluster"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var clustersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List clusters",
	Run: func(cmd *cobra.Command, args []string) {
		clusterSdk := cluster.New(conf)
		clusters, err := clusterSdk.List(&cluster.ListClusterInput{})
		if err != nil {
			log.Fatal(err)
		}

		printer.Print(clusters)
	},
}

func init() {
	clustersCmd.AddCommand(clustersListCmd)
}
