package client

type Addon struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`

	OciRegistry      string `json:"oci_registry"`
	OciVersion       string `json:"oci_version,omitempty"`
	PullPolicy       string `json:"pull_policy"`
	ActivationPolicy string `json:"activation_policy"`

	Tags string `json:"tags,omitempty"`

	Timestamps
}

type AddonSubscription struct {
	Name string `json:"name"`
	Timestamps
}

type CreateAddonInput struct {
	Name             string `json:"name"`
	OciRegistry      string `json:"oci_registry"`
	OciVersion       string `json:"oci_version,omitempty"`
	PullPolicy       string `json:"pull_policy,omitempty"`
	ActivationPolicy string `json:"activation_policy,omitempty"`
	Template         string `json:"template,omitempty"`
}

// :name, :oci_registry, :oci_version, :pull_policy, :activation_policy, :template

type createAddonInput struct {
	Addon *CreateAddonInput `json:"addon"`
}

func (c *Client) ListAddons() (*[]Addon, error) {
	addons := new([]Addon)
	resp, err := c.Handler.
		Get("addons").
		ReceiveSuccess(addons)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return addons, err
}

func (c *Client) GetAddon(addonId string) (*Addon, error) {
	addon := new(Addon)
	resp, err := c.Handler.
		Get("addons/" + addonId).
		ReceiveSuccess(addon)
	if apiError := c.isApiError(resp); apiError != nil {
		err = apiError
	}
	return addon, err
}
