package manage

import (
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/dbclient"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
)

type Manage struct {
	pd.Lister
	Incidents map[string][]dbclient.Incident
}
