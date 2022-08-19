package utils

func JsonError(err error) string {
	return "{\"error\":\"" + err.Error() + "\"}"
}

func JsonMessage(mess string) string {
	return "{\"message\":\"" + mess + "\"}"
}

type Message struct {
	Message string `json:"message"`
}

type Error struct {
	Error string `json:"error"`
}
