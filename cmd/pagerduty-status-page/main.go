package main

import (
	"fmt"
	"log"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/dbclient"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/ui"
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	PagerDutyAuthToken string `required:"true" split_words:"true"`
}

func main() {
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	pdconn := pagerduty.NewClient(s.PagerDutyAuthToken)
	pdclient := pd.NewAPIClient(pdconn)
	incidents := make(map[string][]dbclient.Incident)
	manager := manage.Manage{pdclient, incidents}
	u := ui.Ui{pdclient, incidents}
	err = manager.GetIncidents()
	if err != nil {
		fmt.Println("Problem")
	}
	u.Listen()
}
