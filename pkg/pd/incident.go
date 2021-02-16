package pd

import (
	"github.com/PagerDuty/go-pagerduty"
)

func (c *APIClient) ListI(opts pagerduty.ListIncidentsOptions) (*pagerduty.ListIncidentsResponse, error) {
	eps, err := c.client.ListIncidents(opts)
	return eps, err
}
