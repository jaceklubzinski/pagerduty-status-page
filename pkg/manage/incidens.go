package manage

import (
	"regexp"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	log "github.com/sirupsen/logrus"
)

func (m *Manage) FilterResolvedIncidents(lastUpdate string, alertRegex string) ([]Incident, error) {
	var c uint = 25
	//e := `^\[FIRING:.\](\w+)\s`
	var inc []Incident

	r := regexp.MustCompile(alertRegex)

	for {
		opts := pagerduty.ListIncidentsOptions{
			APIListObject: pagerduty.APIListObject{
				Offset: c,
				Limit:  25,
			},
			Until:    time.Now().String(),
			Since:    lastUpdate,
			Statuses: []string{"resolved"},
		}
		list, err := m.Lister.ListI(opts)
		if err != nil {
			return inc, err
		}
		for _, p := range list.Incidents {
			if result := r.FindStringSubmatch(p.Title); result != nil {
				i := Incident{
					Name:         result[1],
					Service:      p.Service.Summary,
					Urgency:      p.Urgency,
					CreatedAt:    createdAgo(p.CreatedAt),
					LastChangeAt: toDate(p.LastStatusChangeAt),
					Duration:     timeDiff(p.CreatedAt, p.LastStatusChangeAt),
					Team:         p.Teams[0].Summary,
					PDLink:       p.HTMLURL,
					Status:       p.Status,
				}
				inc = append(inc, i)
			}
		}
		c = c + 25
		if !list.More {
			break
		}
	}
	log.Debug("Resolved incident retrived from PagerDuty")
	return inc, nil
}

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
			PDLink:    p.HTMLURL,
			Status:    p.Status,
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
