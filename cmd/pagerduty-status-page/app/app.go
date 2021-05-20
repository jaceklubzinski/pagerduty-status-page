package app

import (
	"time"

	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	log "github.com/sirupsen/logrus"
)

func Run(manager manage.Manage) {

	go func() {
		ticker := time.NewTicker(300 * time.Second)

		for ; true; <-ticker.C {
			manager.ClearIncidents()
			err := manager.GetServices()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the service list")
			}
			err = manager.GetIncidents()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the incident list")
			}
		}
	}()

}
