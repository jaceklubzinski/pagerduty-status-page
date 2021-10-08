package manage

import (
	"html/template"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func (u *Manage) serviceTotalAlerts(name string) int {
	type total struct {
		Service string
		Total   int
	}
	var t []total
	u.DB.Model(Incident{}).Select("service,count(*) as total").Where("status = ?", "resolved").Group("service").Having("service = ?", name).Find(&t)
	return t[0].Total
}
func (u *Manage) serviceUrgentHighCount(name string) int {
	type urgentHigh struct {
		Service string
		High    int
	}
	var highCount []urgentHigh
	var count int

	u.DB.Model(Incident{}).Select("service,count(*) as High").Where("status = ?", "resolved").Group("service,urgency").Having("urgency = ? and service = ?", "high", name).Find(&highCount)

	if len(highCount) == 0 {
		count = 0
	} else {
		count = highCount[0].High
	}
	return count
}
func (u *Manage) serviceUrgentLowCount(name string) int {
	type urgentLow struct {
		Service string
		Low     int
	}

	var lowCount []urgentLow
	var count int

	u.DB.Model(Incident{}).Select("service,count(*) as Low").Where("status = ?", "resolved").Group("service,urgency").Having("urgency = ? and service = ?", "low", name).Find(&lowCount)
	if len(lowCount) == 0 {
		count = 0
	} else {
		count = lowCount[0].Low
	}
	return count
}
func (u *Manage) alertByService(name string) map[string]map[string]int {
	type a struct {
		Name    string
		Urgency string
		Count   int
	}
	var aa []a

	u.DB.Model(Incident{}).Select("name,urgency,count(*) as count").Where("status = ?", "resolved").Where("service = ?", name).Group("name,urgency").Order("3 desc").Find(&aa)

	alerts := make(map[string]map[string]int)
	for _, alert := range aa {
		if alerts[alert.Name] == nil {
			alerts[alert.Name] = map[string]int{}
		}
		alerts[alert.Name][alert.Urgency] = alert.Count
	}

	return alerts
}
func (u *Manage) alertUrgentHighCount(name string) int {
	type urgentHigh struct {
		Service string
		High    int
	}
	var highCount []urgentHigh
	var count int

	u.DB.Model(Incident{}).Select("name,count(*) as High").Where("status = ?", "resolved").Group("name,urgency").Having("urgency = ? and name = ?", "high", name).Find(&highCount)

	if len(highCount) == 0 {
		count = 0
	} else {
		count = highCount[0].High
	}
	return count
}
func (u *Manage) alertUrgentLowCount(name string) int {
	type urgentLow struct {
		Service string
		Low     int
	}

	var lowCount []urgentLow
	var count int

	u.DB.Model(Incident{}).Select("name,count(*) as Low").Where("status = ?", "resolved").Group("name,urgency").Having("urgency = ? and name = ?", "low", name).Find(&lowCount)
	if len(lowCount) == 0 {
		count = 0
	} else {
		count = lowCount[0].Low
	}
	return count
}

func (u *Manage) alertDuration(name string) time.Duration {
	type d struct {
		Avg float64
	}

	var dd []d

	u.DB.Model(Incident{}).Select("avg(duration) as avg").Where("name = ? and status = ?", name, "resolved").Find(&dd)
	d2 := time.Duration(int(dd[0].Avg)) * time.Second

	return d2
}

func (u *Manage) services(w http.ResponseWriter, req *http.Request) {
	fmap := template.FuncMap{
		"trim":                   u.trimService,
		"serviceTotalAlerts":     u.serviceTotalAlerts,
		"serviceUrgentHighCount": u.serviceUrgentHighCount,
		"serviceUrgentLowCount":  u.serviceUrgentLowCount,
		"alertByService":         u.alertByService,
		"alertDuration":          u.alertDuration,
	}

	t := template.Must(template.New("services.tmpl").Funcs(fmap).ParseFiles("services.tmpl", "navbar.tmpl"))

	type s struct {
		Service string
		Count   int
	}
	var ss []s
	u.DB.Model(Incident{}).Select("service, count(*) as count").Where("status = ?", "resolved").Group("service").Order("2 desc").Find(&ss)

	err := t.Execute(w, ss)
	if err != nil {
		log.Errorln(err)
	}
}

func (u *Manage) incidents(w http.ResponseWriter, req *http.Request) {
	fmap := template.FuncMap{
		"alertUrgentLowCount":  u.alertUrgentLowCount,
		"alertUrgentHighCount": u.alertUrgentHighCount,
		"alertDuration":        u.alertDuration,
	}

	t := template.Must(template.New("incidents.tmpl").Funcs(fmap).ParseFiles("incidents.tmpl", "navbar.tmpl"))

	type result struct {
		Name  string
		Count int
	}
	var r []result
	u.DB.Model(Incident{}).Select("name,count(*) as count").Group("name").Order("2 desc").Find(&r)

	err := t.Execute(w, r)
	if err != nil {
		log.Errorln(err)
	}
}
