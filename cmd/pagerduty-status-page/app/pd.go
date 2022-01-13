package app

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
)

func NewPDStatus() manage.Manage {
	initLogs()
	db := initDB()
	s := initENV()

	pdconn := pagerduty.NewClient(s.PagerDutyAuthToken)
	pdclient := pd.NewAPIClient(pdconn)
	incidents := make(map[string]map[string][]manage.Incident)
	manager := manage.Manage{pdclient, incidents, db}
	return manager
}
