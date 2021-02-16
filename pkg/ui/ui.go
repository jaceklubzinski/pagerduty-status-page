package ui

import (
	"html/template"
	"net/http"

	"github.com/jaceklubzinski/pagerduty-status-page/pkg/dbclient"
	"github.com/jaceklubzinski/pagerduty-status-page/pkg/pd"
	log "github.com/sirupsen/logrus"
)

type Ui struct {
	pd.Lister
	dbclient.Storer
}

func (u *Ui) Listen() {
	http.HandleFunc("/", u.manage)
	log.Info(http.ListenAndServe(":9090", nil))
}

func (u *Ui) manage(w http.ResponseWriter, req *http.Request) {

	s, err := u.Storer.GetServices()
	if err != nil {
		log.Errorln(err)
	}

	UITemplate, err := template.New("incidents.tmpl").Funcs(template.FuncMap{
		"getIncidents": func(service string) []dbclient.Incident {
			i, err := u.GetIncidents(service)
			if err != nil {
				log.Errorln(err)
			}
			return i
		},
	}).ParseFiles("incidents.tmpl")
	if err != nil {
		log.Errorln(err)
	}

	err = UITemplate.Execute(w, s)

	//t := template.Must(template.ParseFiles("pkg/ui/services_v3.tmpl"))

	//err = t.Execute(w, s)
	if err != nil {
		log.Errorln(err)
	}

}
