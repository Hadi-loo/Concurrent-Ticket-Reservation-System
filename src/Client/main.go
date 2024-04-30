package main

import (
	"net/http"
)

const (
	LOGS_INFO_PREFIX  = "\033[33m[INFO]\033[0m "
	LOGS_WARN_PREFIX  = "\033[31m[WARN]\033[0m "
	LOGS_ERROR_PREFIX = "\033[31m[ERROR]\033[0m "
)

const (
	ListEventsUrl  = "http://localhost:8080/events"
	CreateEventUrl = "http://localhost:8080/events/create"
	BookTicketsUrl = "http://localhost:8080/events/book"
)

func main() {
	client := http.Client{}
	SetupClient(&client)
	GetInput(&client)
}
