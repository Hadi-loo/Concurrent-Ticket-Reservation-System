package main

import (
	"Ticket_Resevation/src/TicketService"
	"net/http"
)

const (
	LOGS_INFO_PREFIX  = "\033[33m[INFO]\033[0m "
	LOGS_WARN_PREFIX  = "\033[31m[WARN]\033[0m "
	LOGS_ERROR_PREFIX = "\033[31m[ERROR]\033[0m "
)

func main() {
	ticketService := TicketService.TicketService{}
	TicketService.InitTicketService("database/", "events.json", &ticketService)
	defer TicketService.Save("database/", "events.json", &ticketService)

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		ListEventsHandler(w, r, &ticketService)
	})
	http.HandleFunc("/events/create", func(w http.ResponseWriter, r *http.Request) {
		CreateEventHandler(w, r, &ticketService)
	})
	http.HandleFunc("/events/book", func(w http.ResponseWriter, r *http.Request) {
		BookTicketsHandler(w, r, &ticketService)
	})

	startServer()

}
