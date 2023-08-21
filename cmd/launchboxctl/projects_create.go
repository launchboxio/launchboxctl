package main

import (
	"github.com/launchboxio/launchboxctl/pkg/client"
	"github.com/spf13/cobra"
)

var projectsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		input := &client.CreateProjectInput{
			Name: name,
		}

		if memory, _ := cmd.Flags().GetInt("memory"); memory != 0 {
			input.Memory = memory
		}
		if cpu, _ := cmd.Flags().GetInt("cpu"); cpu != 0 {
			input.Cpu = cpu
		}
		if disk, _ := cmd.Flags().GetInt("disk"); disk != 0 {
			input.Disk = disk
		}
		if gpu, _ := cmd.Flags().GetInt("gpu"); gpu != 0 {
			input.Gpu = gpu
		}

		project, err := sdk.CreateProject(input)
		if err != nil {
			return err
		}

		return printer.Output(project)
	},
}

func init() {
	projectsCreateCmd.Flags().Int("cpu", 0, "The number of CPU cores")
	projectsCreateCmd.Flags().Int("memory", 0, "The amount of memory, in GB")
	projectsCreateCmd.Flags().Int("disk", 0, "Amount of disk space, in GB")
	projectsCreateCmd.Flags().String("kubernetes-version", "", "Project Kubernetes API Version")
	projectsCreateCmd.Flags().String("name", "", "Project name")
	projectsCreateCmd.Flags().Int("gpu", 0, "The number of GPUs")

	projectsCreateCmd.MarkFlagRequired("name")

	projectsCmd.AddCommand(projectsCreateCmd)
}
