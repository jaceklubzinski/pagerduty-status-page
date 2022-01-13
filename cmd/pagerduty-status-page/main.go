package main

import "github.com/jaceklubzinski/pagerduty-status-page/cmd/pagerduty-status-page/app"

func main() {
	manager := app.NewPDStatus()
	app.RunIncidents(manager)
	app.RunAnalytics(manager)
	manager.Listen()
}
