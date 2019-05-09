package clients

import "github.com/hashicorp/consul/api"

// Consul defines an interface for a Consul client
type Consul interface {
	// GetIntentions returns a list of intentions currently configured in
	// Consul
	GetIntentions() ([]*api.Intention, error)

	// DeleteIntention deletes an intention in Consul
	DeleteIntention() error

	// CreateIntention creates an intention in Consul
	CreateIntention(source string, destination string) (bool, error)
}

// ConsulImpl concrete implementation of the Consul client interface
type ConsulImpl struct {
	client *api.Client
}

// NewConsul creates a new Consul client
func NewConsul(httpAddr, aclToken string) (Consul, error) {
	conf := api.DefaultConfig()
	conf.Address = httpAddr
	conf.Token = aclToken

	cli, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &ConsulImpl{cli}, nil
}

// GetIntentions returns a list of intentions currently configured in
// Consul
func (c *ConsulImpl) GetIntentions() ([]*api.Intention, error) {
	var intentions []*api.Intention
	intentions, _, err := c.client.Connect().Intentions(nil)
	if err != nil {
		return intentions, err
	}

	return intentions, nil
}

// CreateIntention creates an intention in Consul
func (c *ConsulImpl) CreateIntention(source string, destination string) (bool, error) {
	in := api.Intention{
		SourceName:      source,
		DestinationName: destination,
		Action:          api.IntentionActionAllow,
		Description:     "Automatically added by Kubernetes",
	}

	_, _, err := c.client.Connect().IntentionCreate(&in, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteIntention deletes an intention in Consul
func (c *ConsulImpl) DeleteIntention() error {
	return nil
}
