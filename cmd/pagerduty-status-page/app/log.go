package app

import log "github.com/sirupsen/logrus"

func initLogs() {
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
	log.SetFormatter(formatter)
	log.SetLevel(log.DebugLevel)
}
