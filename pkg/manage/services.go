package manage

import (
	log "github.com/sirupsen/logrus"
)

func (m *Manage) GetServices() error {
	services, err := m.Lister.ListS()
	if err != nil {
		return err
	}
	for _, s := range services {
		m.Incidents[s.Name] = make(map[string][]Incident)
	}
	log.Debug("services retrived from PagerDuty")

	return nil
}
