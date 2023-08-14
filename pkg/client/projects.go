package client

type Project struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status,omitempty"`
	Slug   string `json:"slug"`

	Memory int `json:"memory"`
	Cpu    int `json:"cpu"`
	Disk   int `json:"disk"`
	Gpu    int `json:"gpu"`

	Timestamps
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
