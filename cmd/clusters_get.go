package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/cluster"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var clustersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information for a specific cluster",
	Run: func(cmd *cobra.Command, args []string) {
		clusterId, _ := cmd.Flags().GetInt("cluster-id")
		clusterSdk := cluster.New(conf)
		cluster, err := clusterSdk.Get(&cluster.GetClusterInput{
			ClusterId: clusterId,
		})
		if err != nil {
			log.Fatal(err)
		}

		printer.Print(cluster)
	},
}

func init() {
	clustersCmd.AddCommand(clustersGetCmd)

	clustersGetCmd.Flags().Int("cluster-id", 0, "ID of the cluster to get")
	_ = clustersGetCmd.MarkFlagRequired("cluster-id")
}
