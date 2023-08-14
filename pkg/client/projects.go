package client

import (
	"fmt"
	"io"
)

type Project struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Status string `json:"status,omitempty"`
	Slug   string `json:"slug,omitempt"`

	Memory int `json:"memory,omitempty"`
	Cpu    int `json:"cpu,omitempty"`
	Disk   int `json:"disk,omitempty"`
	Gpu    int `json:"gpu,omitempty"`

	Timestamps
}

type CreateProjectInput struct {
	Name   string `json:"name"`
	Memory int    `json:"memory,omitempty"`
	Cpu    int    `json:"cpu,omitempty"`
	Disk   int    `json:"disk,omitempty"`
	Gpu    int    `json:"gpu,omitempty"`
}

type createInput struct {
	Project *CreateProjectInput `json:"project"`
}

func (c *Client) ListProjects() (*[]Project, error) {
	projects := new([]Project)
	_, err := c.Handler.
		Get("projects").
		ReceiveSuccess(projects)
	return projects, err
}

func (c *Client) GetProject(projectId string) (*Project, error) {
	project := new(Project)
	_, err := c.Handler.
		Get("projects/" + projectId).
		ReceiveSuccess(project)
	return project, err
}

func (c *Client) CreateProject(input *CreateProjectInput) (*Project, error) {
	project := new(Project)
	resp, err := c.Handler.
		Post("projects").
		BodyJSON(createInput{input}).
		ReceiveSuccess(project)
	fmt.Println(resp)
	bytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	return project, err
}
