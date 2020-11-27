package log

import log "github.com/sirupsen/logrus"

func DefaultLogger() *log.Logger {
	return log.New()
}
