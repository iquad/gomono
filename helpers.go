package gomono

import (
	"os"
	"runtime/debug"
)

func RecoverMe() {
	if rec := recover(); rec != nil {
		logger.Printf("%v", rec)

		errorCounter++

		if errorCounter > 10 {
			logger.Print("maximum error handling exceed, going shutdown.")
			debug.PrintStack()
			os.Exit(1)
		}
	}
}
