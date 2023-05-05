package cmd

import "github.com/spf13/cobra"

type Project struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug,omitempty"`

	ClusterId int          `json:"cluster_id"`
	Owner     ProjectOwner `json:"owner"`

	Memory int `json:"memory,omitempty"`
	Cpu    int `json:"cpu,omitempty"`
	Disk   int `json:"disk,omitempty"`

	Status string   `json:"status,omitempty"`
	Tags   []string `json:"tag_list,omitempty"`
	//Addons []
}

type ProjectOwner struct {
	Email string `json:"email"`
}

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Interact with projects",
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	projectsCmd.AddCommand(projectsListCmd)
	projectsCmd.AddCommand(projectsCreateCmd)
}
