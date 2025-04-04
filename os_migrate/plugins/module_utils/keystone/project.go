package keystone

import "context"
import "fmt"
import "github.com/gophercloud/gophercloud/v2"
import "github.com/gophercloud/gophercloud/v2/openstack"
import "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"


type ProjectOpts struct {
	ID string `json:"id"`
	Name string `json:"name,omitempty"`
	DomainID string `json:"domain_id,omitempty"`
	IsDomain bool `json:"is_domain,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Description string `json:"description,omitempty"`
}

func ListProjects(client *gophercloud.ProviderClient, project ProjectOpts) (bool, error) {
	fmt.Printf("project: %+v\n", project)
	return true, nil
}