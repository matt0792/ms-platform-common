package scheduleclient

import "github.com/matt0792/mscommon/s2s"

type ScheduleClient struct {
	s2s.Transport
}

func NewScheduleClient(baseUrl, token string) *ScheduleClient {
	return &ScheduleClient{
		Transport: s2s.Transport{
			BaseUrl: baseUrl,
			Token:   token,
		},
	}
}
