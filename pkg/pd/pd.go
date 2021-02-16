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

type Lister interface {
	ListI(opts pagerduty.ListIncidentsOptions) (*pagerduty.ListIncidentsResponse, error)
	ListS() ([]pagerduty.Service, error)
}
