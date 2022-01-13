package app

import (
	"time"

	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Specification struct {
	PagerDutyAuthToken string `required:"true" split_words:"true"`
	ResolvedAlertRegex string `required:"true" split_words:"true"`
}

func initENV() Specification {
	var s Specification
	err := envconfig.Process("PDSTATUS", &s)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Warn("problem with ENVIRONMENT variables")
	}
	return s
}

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("incidents.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatalln("Cant initialize sqlite database")
	}

	err = db.AutoMigrate(&manage.Incident{})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatalln("Problem with database migration")
	}
	return db
}

func RunIncidents(manager manage.Manage) {
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

func RunAnalytics(manager manage.Manage) {
	s := initENV()

	go func() {
		ticker := time.NewTicker(60 * time.Minute)

		for ; true; <-ticker.C {
			var incidents []manage.Incident

			lastUpdate, err := manager.GetLastUpdateDaysAgo()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the last update date from DB")
			}
			incidents, err = manager.FilterResolvedIncidents(lastUpdate, s.ResolvedAlertRegex)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the resolved incident list from PagerDuty")
			}

			for _, i := range incidents {
				manager.DB.Create(&manage.Incident{
					Name:         i.Name,
					Service:      i.Service,
					Status:       i.Status,
					Urgency:      i.Urgency,
					Duration:     i.Duration,
					LastChangeAt: i.LastChangeAt,
				})
			}

			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Warn("problem with retrieving the resolved incident list from database")
			}
		}
	}()
}
