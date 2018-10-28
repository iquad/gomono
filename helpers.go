package gomono

import (
	"runtime/debug"
)

func RecoverMe() {
	if rec := recover(); rec != nil {
		logger.Printf("%v", rec)
		debug.PrintStack()

		errorCounter++

		if errorCounter > 10 {
			logger.Fatal("maximum error handling exceed, going shutdown.")
		}
	}
}
