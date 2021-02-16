package dbclient

import (
	"database/sql"
	"log"

	"github.com/PagerDuty/go-pagerduty"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/pdstatus")
}

func RunDB() *Store {

	dbClient, err := connectDB()
	if err != nil {
		log.Fatalf("Can't create database connection: %v", err)
	}
	db := newDB(dbClient)
	err = db.createTable()
	if err != nil {
		log.Fatalf("Can't create status table: %v", err)
	}
	return db
}

func (d *Store) createTable() error {
	s := `
	CREATE TABLE IF NOT EXISTS status(
			name TEXT,
			service TEXT,
			status TEXT,
			urgency TEXT,
			createdAt TEXT
	);
	`
	_, err := d.db.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func (d *Store) deleteAll() error {
	s := `
	DELETE FROM status;
	`
	_, err := d.db.Exec(s)
	if err != nil {
		return err
	}
	return nil
}

func (d *Store) deleteRow(title string) error {
	stmt, err := d.db.Prepare("DELETE FROM status where title=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

func (d *Store) AddRow(i *pagerduty.Incident) error {
	stmt, err := d.db.Prepare("REPLACE INTO status(name,service,status,urgency,createdAt) values(?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(i.Title, i.Service.Summary, i.Status, i.Urgency, i.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (d *Store) GetServices() ([]string, error) {
	var services []string
	r, err := d.db.Query("select distinct service from status;")
	if err != nil {
		return services, err
	}
	for r.Next() {
		var s string
		err := r.Scan(&s)
		if err != nil {
			return services, err
		}
		services = append(services, s)
	}
	return services, nil
}

func (d *Store) GetIncidents(service string) ([]Incident, error) {
	var incidents []Incident
	r, err := d.db.Query("select * from status where service = ?", service)
	if err != nil {
		return incidents, err
	}
	for r.Next() {
		var i Incident
		err := r.Scan(&i.Name, &i.Service, &i.Status, &i.Urgency, &i.Createdat)
		if err != nil {
			return incidents, err
		}
		incidents = append(incidents, i)
	}
	return incidents, nil
}
