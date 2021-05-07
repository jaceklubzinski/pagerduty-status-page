package pd

import (
	"github.com/PagerDuty/go-pagerduty"
)

//List list all PagerDuty services
func (c *APIClient) ListS() ([]pagerduty.Service, error) {
	opts := pagerduty.ListServiceOptions{
		APIListObject: pagerduty.APIListObject{
			Limit: 100,
		},
	}
	s, err := c.client.ListServices(opts)
	return s.Services, err
}
