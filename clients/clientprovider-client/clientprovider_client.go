package clientproviderclient

import (
	"github.com/matt0792/mscommon/commonmodels"
	"github.com/matt0792/mscommon/s2s"
)

type ClientProviderClient struct {
	s2s.Transport
}

func NewClientProviderClient(baseUrl, token string) *ClientProviderClient {
	return &ClientProviderClient{
		Transport: s2s.Transport{
			BaseUrl: baseUrl,
			Token:   token,
		},
	}
}

func (c *ClientProviderClient) GetServiceLocation(serviceName string) (string, error) {
	var resp commonmodels.Response[string]
	err := c.Call("GET", "/services/"+serviceName+"/location", nil, &resp)
	return s2s.UnwrapResponse(resp, err)
}

func (c *ClientProviderClient) SetServiceConfig(config commonmodels.MicroserviceConfig) (string, error) {
	var resp commonmodels.Response[string]
	err := c.Call("POST", "/services/register", config, &resp)
	return s2s.UnwrapResponse(resp, err)
}

func (c *ClientProviderClient) GetAllServices() ([]commonmodels.MicroserviceConfig, error) {
	var resp commonmodels.Response[[]commonmodels.MicroserviceConfig]
	err := c.Call("POST", "/service/register", nil, &resp)
	return s2s.UnwrapResponse(resp, err)
}
