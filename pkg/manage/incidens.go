package manage

import (
	"github.com/PagerDuty/go-pagerduty"
	log "github.com/sirupsen/logrus"
)

//GetIncidents get and format PagerDuty incidents
func (m *Manage) GetIncidents() error {
	opts := pagerduty.ListIncidentsOptions{
		APIListObject: pagerduty.APIListObject{
			Limit: 100,
		},
		Statuses: []string{"triggered", "acknowledged"}}
	list, err := m.Lister.ListI(opts)
	if err != nil {
		return err
	}
	for _, p := range list.Incidents {

		i := Incident{
			Name:      p.Title,
			Service:   p.Service.Summary,
			Urgency:   p.Urgency,
			Assigne:   p.Assignments[0].Assignee.Summary,
			CreatedAt: createdAgo(p.CreatedAt),
			Team:      p.Teams[0].Summary,
		}

		m.Incidents[p.Service.Summary][p.Urgency] = append(m.Incidents[p.Service.Summary][p.Urgency], i)
	}
	log.Debug("Incident retrived from PagerDuty")
	return nil
}

//ClearIncidents clear incidents map
func (m *Manage) ClearIncidents() {
	for k := range m.Incidents {
		delete(m.Incidents, k)
	}
}
