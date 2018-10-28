package gomono

func RecoverMe() {
	if rec := recover(); rec != nil {
		logger.Printf("%v", rec)

		errorCounter++

		if errorCounter > 10 {
			logger.Fatalf("maximum error handling exceed, going shutdown.")
		}
	}
}
