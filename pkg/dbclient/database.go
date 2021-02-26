package dbclient

import (
	"database/sql"

	"github.com/PagerDuty/go-pagerduty"
)

type Store struct {
	db        *sql.DB
	Incidents []Incident
}

type Incident struct {
	Name      string
	Service   string
	Status    string
	Urgency   string
	Createdat string
}

type Incidents struct {
	Incident []Incident
}

type Storer interface {
	createTable() error
	deleteAll() error
	deleteRow(title string) error
	AddRow(i *pagerduty.Incident) error
	GetServices() ([]string, error)
	GetIncidents(service string) ([]Incident, error)
}

func newDB(db *sql.DB) *Store {
	return &Store{db: db}
}
