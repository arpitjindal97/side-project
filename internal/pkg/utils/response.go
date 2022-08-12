package utils

func JsonError(err error) string {
	return "{\"error\":\"" + err.Error() + "\"}"
}

func JsonMessage(mess string) string {
	return "{\"message\":\"" + mess + "\"}"
}
