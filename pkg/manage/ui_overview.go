package manage

import (
	"html/template"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (u *Manage) Listen() {
	http.HandleFunc("/", u.manage)
	http.HandleFunc("/services", u.services)
	http.HandleFunc("/incidents", u.incidents)

	log.Info(http.ListenAndServe(":9090", nil))
}

func (u *Manage) manage(w http.ResponseWriter, req *http.Request) {
	fmap := template.FuncMap{
		"trim": u.trimService,
	}

	t := template.Must(template.New("overview.tmpl").Funcs(fmap).ParseFiles("pkg/manage/templates/overview.tmpl", "pkg/manage/templates/navbar.tmpl", "pkg/manage/templates/searchbar.tmpl", "pkg/manage/templates/style.tmpl"))

	err := t.Execute(w, u.Incidents)
	if err != nil {
		log.Errorln(err)
	}
}

func (u *Manage) trimService(name string) string {
	return (strings.ReplaceAll(name, " ", ""))
}
