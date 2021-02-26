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
	Incidents map[string][]dbclient.Incident
}

func (u *Ui) Listen() {
	http.HandleFunc("/", u.manage)
	log.Info(http.ListenAndServe(":9090", nil))
}

func (u *Ui) manage(w http.ResponseWriter, req *http.Request) {

	t := template.Must(template.ParseFiles("pkg/ui/services_v3.tmpl"))

	err := t.Execute(w, u.Incidents)
	if err != nil {
		log.Errorln(err)
	}

}
