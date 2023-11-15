package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/cluster"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var clustersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new cluster",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		domain, _ := cmd.Flags().GetString("domain")

		clusterSdk := cluster.New(conf)
		result, err := clusterSdk.Create(&cluster.CreateClusterInput{
			Domain: domain,
			Name:   name,
		})
		if err != nil {
			log.Fatal(err)
		}

		printer.Print(result.Cluster)
	},
}

func init() {
	clustersCreateCmd.Flags().String("name", "", "The name of the cluster")
	clustersCreateCmd.Flags().String("domain", "", "The domain for project ingress")

	_ = clustersCreateCmd.MarkFlagRequired("name")
	_ = clustersCreateCmd.MarkFlagRequired("domain")
	clustersCmd.AddCommand(clustersCreateCmd)
}
