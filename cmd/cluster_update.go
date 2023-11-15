package cmd

import (
  "github.com/launchboxio/launchbox-go-sdk/service/cluster"
  "github.com/launchboxio/launchboxctl/internal/printer"
  "github.com/spf13/cobra"
  "log"
)

var clustersUpdateCmd = &cobra.Command{
  Use:   "update",
  Short: "Update a clusters details",
  Run: func(cmd *cobra.Command, args []string) {
    domain, _ := cmd.Flags().GetString("domain")
    status, _ := cmd.Flags().GetString("status")
    clusterId, _ := cmd.Flags().GetInt("cluster-id")

    clusterSdk := cluster.New(conf)
    result, err := clusterSdk.Update(&cluster.UpdateClusterInput{
      Domain:    domain,
      Status:    status,
      ClusterId: clusterId,
    })
    if err != nil {
      log.Fatal(err)
    }

    printer.Print(result.Cluster)
  },
}

func init() {
  clustersUpdateCmd.Flags().String("domain", "", "The domain for project ingress")
  clustersUpdateCmd.Flags().String("status", "", "The status for the cluster")
  clustersUpdateCmd.Flags().Int("cluster-id", 0, "The Cluster ID")
  _ = clustersUpdateCmd.MarkFlagRequired("cluster-id")

  clustersCmd.AddCommand(clustersUpdateCmd)
}
