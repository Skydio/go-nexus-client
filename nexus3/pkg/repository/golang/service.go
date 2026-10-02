package golang

import (
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/client"
	"github.com/datadrivers/go-nexus-client/nexus3/pkg/repository/common"
	"github.com/datadrivers/go-nexus-client/nexus3/schema/repository"
)

const (
	goAPIEndpoint = common.RepositoryAPIEndpoint + "/go"
)

type (
	RepositoryGoGroupService  = common.RepositoryService[repository.GoGroupRepository]
	RepositoryGoHostedService = common.RepositoryService[repository.GoHostedRepository]
	RepositoryGoProxyService  = common.RepositoryService[repository.GoProxyRepository]
)

type RepositoryGoService struct {
	client *client.Client

	Group  *RepositoryGoGroupService
	Hosted *RepositoryGoHostedService
	Proxy  *RepositoryGoProxyService
}

func NewRepositoryGoService(c *client.Client) *RepositoryGoService {
	return &RepositoryGoService{
		client: c,

		Group:  NewRepositoryGoGroupService(c),
		Hosted: NewRepositoryGoHostedService(c),
		Proxy:  NewRepositoryGoProxyService(c),
	}
}

func NewRepositoryGoGroupService(c *client.Client) *RepositoryGoGroupService {
	return common.NewRepositoryService[repository.GoGroupRepository](goAPIEndpoint+"/group", c)
}

func NewRepositoryGoHostedService(c *client.Client) *RepositoryGoHostedService {
	return common.NewRepositoryService[repository.GoHostedRepository](goAPIEndpoint+"/hosted", c)
}

func NewRepositoryGoProxyService(c *client.Client) *RepositoryGoProxyService {
	return common.NewRepositoryService[repository.GoProxyRepository](goAPIEndpoint+"/proxy", c)
}
