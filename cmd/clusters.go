package cmd

import "github.com/spf13/cobra"

type Cluster struct {
	Name string `json:"name"`
}

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Interact with clusters",
}

func init() {
	rootCmd.AddCommand(clustersCmd)

	clustersCmd.AddCommand(clustersListCmd)
}
