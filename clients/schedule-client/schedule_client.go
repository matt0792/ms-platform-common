package scheduleclient

import (
	schedulemodels "github.com/matt0792/ms-platform-common/models/schedule-models"
	"github.com/matt0792/mscommon/commonmodels"
	"github.com/matt0792/mscommon/s2s"
)

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

func (c *ScheduleClient) GetCronJob(id string) (schedulemodels.CronJob, error) {
	var resp commonmodels.Response[schedulemodels.CronJob]
	err := c.Call("GET", "/cron/"+id, nil, &resp)
	return s2s.UnwrapResponse(resp, err)
}

func (c *ScheduleClient) GetAllCronJobs() ([]schedulemodels.CronJob, error) {
	var resp commonmodels.Response[[]schedulemodels.CronJob]
	err := c.Call("GET", "/cron/all", nil, &resp)
	return s2s.UnwrapResponse(resp, err)
}

func (c *ScheduleClient) GetActiveCronJobs() ([]schedulemodels.CronJob, error) {
	var resp commonmodels.Response[[]schedulemodels.CronJob]
	err := c.Call("GET", "/cron/active", nil, &resp)
	return s2s.UnwrapResponse(resp, err)
}

func (c *ScheduleClient) SetCronJob(job schedulemodels.CronJob) (string, error) {
	var resp commonmodels.Response[string]
	err := c.Call("POST", "/cron/create", job, &resp)
	return s2s.UnwrapResponse(resp, err)
}

func (c *ScheduleClient) SetCronJobActive(id string, active bool) (string, error) {
	body := struct {
		IsActive bool `json:"isActive"`
	}{IsActive: active}

	var resp commonmodels.Response[string]
	err := c.Call("PATCH", "/cron/"+id+"/activate", body, &resp)
	return s2s.UnwrapResponse(resp, err)
}
