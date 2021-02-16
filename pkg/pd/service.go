package pd

import (
	"github.com/PagerDuty/go-pagerduty"
)

//List list all PagerDuty services
func (c *APIClient) ListS() ([]pagerduty.Service, error) {
	var opts pagerduty.ListServiceOptions
	s, err := c.client.ListServices(opts)
	return s.Services, err

}
