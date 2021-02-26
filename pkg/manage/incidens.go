package manage

import (
	"fmt"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/dbclient"
)

func (m *Manage) GetIncidents() error {
	opts := pagerduty.ListIncidentsOptions{
		Statuses: []string{"triggered", "acknowledged"}}
	list, err := m.Lister.ListI(opts)
	if err != nil {
		return err
	}
	for _, p := range list.Incidents {
		i := dbclient.Incident{
			Name:      p.Title,
			Service:   p.Service.Summary,
			Status:    p.Status,
			Urgency:   p.Urgency,
			Createdat: p.CreatedAt,
		}
		m.Incidents[p.Service.Summary] = append(m.Incidents[p.Service.Summary], i)
	}
	for k, _ := range m.Incidents {
		for _, v := range m.Incidents[k] {
			fmt.Println("Service: ", k)
			fmt.Println("Incident: ", v)
		}
	}
	return nil
}
