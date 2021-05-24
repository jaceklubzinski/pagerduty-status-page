package manage

import (
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
)

type Manage struct {
	pd.Lister
	Incidents map[string]map[string][]Incident
}

/**
 * @todo Add link do PD
 * @body Additional info with direct link to PD incident
 */
type Incident struct {
	Name      string
	Service   string
	Urgency   string
	Assigne   string
	CreatedAt string
	Team      string
}
