package main

import (
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/ui"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	PagerDutyAuthToken string `required:"true" split_words:"true"`
}

func main() {
	var s Specification
	err := envconfig.Process("PDSTATUS", &s)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("problem with ENVIRONMENT variables")
	}
	pdconn := pagerduty.NewClient(s.PagerDutyAuthToken)
	pdclient := pd.NewAPIClient(pdconn)
	incidents := make(map[string]map[string][]manage.Incident)
	manager := manage.Manage{pdclient, incidents}
	u := ui.Ui{pdclient, incidents}
	err = manager.GetServices()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("problem with retrieving the service list")
	}

	go func() {
		ticker := time.NewTicker(600 * time.Second)
		err = manager.GetIncidents()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("problem with retrieving the incident list")
		}
		for _ = range ticker.C {
			for k := range incidents {
				delete(incidents, k)
			}
			err = manager.GetIncidents()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the incident list")
			}
		}
	}()

	u.Listen()
}
