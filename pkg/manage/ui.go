package manage

import (
	"html/template"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (u *Manage) Listen() {
	http.HandleFunc("/", u.manage)
	log.Info(http.ListenAndServe(":9090", nil))
}

func (u *Manage) manage(w http.ResponseWriter, req *http.Request) {
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
