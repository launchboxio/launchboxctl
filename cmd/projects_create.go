package cmd

import (
	"github.com/launchboxio/launchbox-go-sdk/service/project"
	"github.com/launchboxio/launchboxctl/internal/printer"
	"github.com/spf13/cobra"
	"log"
)

var projectsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")

		memory, _ := cmd.Flags().GetInt("memory")
		cpu, _ := cmd.Flags().GetInt("cpu")
		disk, _ := cmd.Flags().GetInt("disk")

		//tags, _ := cmd.Flags().GetStringArray("tags")

		input := &project.CreateProjectInput{
			Cpu:    cpu,
			Memory: memory,
			Disk:   disk,
			Name:   name,
		}

		projectSdk := project.New(conf)
		out, err := projectSdk.Create(input)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(out)
	},
}

func init() {
	projectsCreateCmd.Flags().String("name", "", "Project name")
	_ = projectsCreateCmd.MarkFlagRequired("name")
	//
	//projectsCreateCmd.Flags().Int("cluster", 0, "ID of the cluster to attach the project to")
	//_ = projectsCreateCmd.MarkFlagRequired("cluster")

	projectsCreateCmd.Flags().Bool("wait", false, "Wait for project creation to finish")
	projectsCreateCmd.Flags().Int("memory", 4096, "Memory for the project (GB)")
	projectsCreateCmd.Flags().Int("cpu", 2, "CPU cores to assign")
	projectsCreateCmd.Flags().Int("disk", 50, "Disk size, in GB")
	projectsCreateCmd.Flags().StringArray("tags", []string{}, "Tags for the project")
}
