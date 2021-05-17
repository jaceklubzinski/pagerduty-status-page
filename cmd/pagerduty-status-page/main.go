package main

import "github.com/jaceklubzinski/pagerduty-status-page/cmd/pagerduty-status-page/app"

func main() {
	manager := app.NewPDStatus()
	app.Run(manager)
	manager.Listen()
}
