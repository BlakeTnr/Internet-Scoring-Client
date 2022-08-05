package internetchecker

func VerifyInternetConnection(attempts int) bool {
	for i := 0; i < attempts; i++ {
		randomWebsite := RandomWebsite()
		respond := randomWebsite.VerifyConnection()

		if respond {
			return true
		}
	}
	return false
}
