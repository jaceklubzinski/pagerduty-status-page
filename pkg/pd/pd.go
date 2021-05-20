package pd

import (
	"github.com/PagerDuty/go-pagerduty"
)

type APIClient struct {
	client *pagerduty.Client
}

func NewAPIClient(client *pagerduty.Client) *APIClient {
	return &APIClient{client: client}
}

/**
 * @todo Add admin on duty
 * @body Incidents can be assigned to the person that is not on duty currently, additional info about admin on duty will be appreciated
 */
type Lister interface {
	ListI(opts pagerduty.ListIncidentsOptions) (*pagerduty.ListIncidentsResponse, error)
	ListS() ([]pagerduty.Service, error)
}
