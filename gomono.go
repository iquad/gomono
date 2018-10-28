package gomono

import "log"

var (
	logger       *log.Logger
	errorCounter = 0
)

func SetLogger(log *log.Logger) {
	logger = log
}
