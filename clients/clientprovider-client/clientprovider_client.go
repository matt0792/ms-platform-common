package clientproviderclient

import (
	"github.com/matt0792/mscommon/commonmodels"
	"github.com/matt0792/mscommon/s2s"
)

type ClientProviderClient struct {
	s2s.Transport
}

func NewUserService(baseUrl, token string) *ClientProviderClient {
	return &ClientProviderClient{
		Transport: s2s.Transport{
			BaseUrl: baseUrl,
			Token:   token,
		},
	}
}

func (c *ClientProviderClient) GetServiceLocation(serviceName string) (string, error) {
	var resp commonmodels.Response[string]
	err := c.Call("GET", "/service/"+serviceName+"/location", nil, &resp)
	return s2s.UnwrapResponse(resp, err)
}
