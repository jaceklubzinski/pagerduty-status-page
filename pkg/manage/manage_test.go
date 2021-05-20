package manage

import (
	"testing"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/stretchr/testify/assert"
)

type mockAPIClient struct{}

type mockLister interface { //nolint:golint,deadcode,unused
	ListI(opts pagerduty.ListIncidentsOptions) (*pagerduty.ListIncidentsResponse, error)
	ListS() ([]pagerduty.Service, error)
}

func (c *mockAPIClient) ListS() ([]pagerduty.Service, error) {
	return []pagerduty.Service{
		pagerduty.Service{
			Name: "TestS",
		},
	}, nil
}

func (c *mockAPIClient) ListI(opts pagerduty.ListIncidentsOptions) (*pagerduty.ListIncidentsResponse, error) {
	return &pagerduty.ListIncidentsResponse{
		Incidents: []pagerduty.Incident{
			pagerduty.Incident{
				Title: "TestT",
				Service: pagerduty.APIObject{
					Summary: "TestS",
				},
				Urgency:   "High",
				CreatedAt: time.Now().AddDate(0, 0, -2).Format("2006-01-02T15:04:05Z"),
				Assignments: []pagerduty.Assignment{
					pagerduty.Assignment{
						Assignee: pagerduty.APIObject{
							Summary: "TestA",
						},
					},
				},
				Teams: []pagerduty.APIObject{
					pagerduty.APIObject{
						Summary: "TestTeam",
					},
				},
			},
		},
	}, nil
}

func TestGetIncidents(t *testing.T) {
	pdclient := &mockAPIClient{}
	incidents := make(map[string]map[string][]Incident)
	manager := Manage{pdclient, incidents}
	err := manager.GetServices()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get services", err)
	}
	err = manager.GetIncidents()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when get incidents", err)
	}
	assert.Equal(t, incidents["TestS"]["High"][0].Name, "TestT")
	assert.Equal(t, incidents["TestS"]["High"][0].Assigne, "TestA")
	assert.Equal(t, incidents["TestS"]["High"][0].CreatedAt, "2 days")
}
