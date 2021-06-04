package manage

import (
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
)

//Manage struct fo managaing PagerDuty incidents
type Manage struct {
	pd.Lister
	Incidents map[string]map[string][]Incident
}

//Incident internal incident struct to map PagerDuty incident
type Incident struct {
	Name      string
	Service   string
	Urgency   string
	Assigne   string
	CreatedAt string
	Team      string
	PDLink    string
}
