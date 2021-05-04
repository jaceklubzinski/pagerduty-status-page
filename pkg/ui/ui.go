package ui

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/jaceklubzinski/pagerduty-status-page/pkg/manage"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
	log "github.com/sirupsen/logrus"
)

type Ui struct {
	pd.Lister
	Incidents map[string]map[string][]manage.Incident
}

func (u *Ui) Listen() {
	http.HandleFunc("/", u.manage)
	log.Info(http.ListenAndServe(":9090", nil))
}

func (u *Ui) manage(w http.ResponseWriter, req *http.Request) {
	/*
		t := template.Must(template.ParseFiles("pkg/ui/services_v4.tmpl"))

		err := t.Execute(w, u.Incidents)
		if err != nil {
			log.Errorln(err)
		}
	*/
	t := template.Must(template.New("incidents.tmpl").Funcs(template.FuncMap{
		"trim": func(name string) string {
			return (strings.ReplaceAll(name, " ", ""))
		},
	}).ParseFiles("incidents.tmpl"))

	err := t.Execute(w, u.Incidents)
	if err != nil {
		log.Errorln(err)
	}
}
