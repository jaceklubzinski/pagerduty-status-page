package manage

import (
	"github.com/PagerDuty/go-pagerduty"
)

func (i *Manage) Incidents() error {
	opts := pagerduty.ListIncidentsOptions{
		Statuses: []string{"triggered", "acknowledged"}}
	list, err := i.Lister.ListI(opts)
	if err != nil {
		return err
	}
	for _, p := range list.Incidents {
		//fmt.Printf("ID: %s Name: %s  Service: %s Body: %s Status: %s\n", p.APIObject.ID, p.Title, p.Service.Summary, p.Body, p.Status)
		i.Storer.AddRow(&p)

	}
	return nil
}
