package app

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	PagerDutyAuthToken string `required:"true" split_words:"true"`
}

func NewPDStatus() manage.Manage {
	InitLogs()
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
	return manager
}
