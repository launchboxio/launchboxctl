package client

import "fmt"

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
	resp, err := c.Handler.
		Get("projects").
		ReceiveSuccess(projects)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return projects, err
}

func (c *Client) GetProject(projectId string) (*Project, error) {
	project := new(Project)
	resp, err := c.Handler.
		Get("projects/" + projectId).
		ReceiveSuccess(project)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return project, err
}

func (c *Client) CreateProject(input *CreateProjectInput) (*Project, error) {
	project := new(Project)
	resp, err := c.Handler.
		Post("projects").
		BodyJSON(createInput{input}).
		ReceiveSuccess(project)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return project, err
}

func (c *Client) DeleteProject(projectId string) error {
	_, err := c.Handler.
		Delete("projects/" + projectId).
		ReceiveSuccess(nil)

	return err
}

type AttachAddonInput struct {
	Name string `json:"name"`
}

type attachAddonInput struct {
	AddonSubscription AttachAddonInput `json:"addon_subscription"`
}

func (c *Client) AttachAddon(projectId string, addonId string, input AttachAddonInput) (*AddonSubscription, error) {
	subscription := new(AddonSubscription)
	resp, err := c.Handler.
		Post("projects/" + projectId + "/addons").
		BodyJSON(attachAddonInput{input}).
		ReceiveSuccess(subscription)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return subscription, err
}

func (c *Client) GetProjectAddons(projectId int) (*[]Addon, error) {
	addons := new([]Addon)
	resp, err := c.Handler.
		Get(fmt.Sprintf("projects/%d/addons", projectId)).
		ReceiveSuccess(addons)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return addons, err
}
