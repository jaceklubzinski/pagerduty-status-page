package app

import (
	"time"

	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	log "github.com/sirupsen/logrus"
)

func Run(manager manage.Manage) {
	err := manager.GetServices()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("problem with retrieving the service list")
	}

	go func() {
		ticker := time.NewTicker(300 * time.Second)
		err = manager.GetIncidents()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("problem with retrieving the incident list")
		}
		for range ticker.C {
			manager.ClearIncidents()
			err = manager.GetIncidents()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the incident list")
			}
		}
	}()

}
