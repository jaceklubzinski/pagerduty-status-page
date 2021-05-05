package manage

import (
	"github.com/PagerDuty/go-pagerduty"
	log "github.com/sirupsen/logrus"
)

func (m *Manage) GetIncidents() error {
	opts := pagerduty.ListIncidentsOptions{
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
		}
		if _, ok := m.Incidents[p.Service.Summary][p.Urgency]; !ok {
			m.Incidents[p.Service.Summary] = make(map[string][]Incident)
		}

		m.Incidents[p.Service.Summary][p.Urgency] = append(m.Incidents[p.Service.Summary][p.Urgency], i)
	}
	log.Debug("Incident retrived from PagerDuty")
	return nil
}
