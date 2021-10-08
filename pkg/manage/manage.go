package manage

import (
	"time"

	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
	"gorm.io/gorm"
)

//Manage struct fo managaing PagerDuty incidents
type Manage struct {
	pd.Lister
	Incidents map[string]map[string][]Incident
	DB        *gorm.DB
}

//Incident internal incident struct to map PagerDuty incident
type Incident struct {
	Name         string
	Service      string
	Urgency      string
	Assigne      string
	CreatedAt    string
	LastChangeAt time.Time
	Duration     int
	Team         string
	PDLink       string
	Status       string
}
