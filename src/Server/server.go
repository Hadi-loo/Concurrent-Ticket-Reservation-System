package main

import (
	"Ticket_Resevation/src/TicketService"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.RWMutex

func startServer() {
	port := ":8080"
	log.SetPrefix(LOGS_INFO_PREFIX)
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func ListEventsHandler(w http.ResponseWriter, r *http.Request, ticketService *TicketService.TicketService) {
	mu.RLock()
	defer mu.RUnlock()

	log.SetPrefix(LOGS_INFO_PREFIX)
	log.Println("List Events Called")
	events := ticketService.ListEvents()

	response, err := json.Marshal(events)
	if err != nil {
		http.Error(w, "Failed to marshal events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request, ticketService *TicketService.TicketService) {
	mu.Lock()
	defer mu.Unlock()

	type CreateEventRequest struct {
		Name         string `json:"name"`
		Date         string `json:"date"`
		TotalTickets string `json:"totalTickets"`
	}

	var req CreateEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		http.Error(w, "Failed to parse date", http.StatusBadRequest)
		return
	}
	totalTickets, err := strconv.Atoi(req.TotalTickets)
	if err != nil {
		http.Error(w, "Failed to parse totalTickets", http.StatusBadRequest)
		return
	}

	event, err := ticketService.CreateEvent(req.Name, date, totalTickets)
	if err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Failed to marshal event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	log.SetPrefix(LOGS_INFO_PREFIX)
	log.Printf("New event \"%s\" with %d tickets created", event.Name, event.TotalTickets)
}

func BookTicketsHandler(w http.ResponseWriter, r *http.Request, ticketService *TicketService.TicketService) {
	mu.Lock()
	defer mu.Unlock()

	type BookTicketsRequest struct {
		ID      string `json:"eventID"`
		Tickets string `json:"numTickets"`
	}

	var req BookTicketsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	tickets, err := strconv.Atoi(req.Tickets)
	if err != nil {
		http.Error(w, "Failed to parse tickets", http.StatusBadRequest)
		return
	}

	ticketIDs, err := ticketService.BookTickets(req.ID, tickets)
	if err != nil {
		http.Error(w, "Failed to book tickets", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(ticketIDs)
	if err != nil {
		http.Error(w, "Failed to marshal ticketIDs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

	log.SetPrefix(LOGS_INFO_PREFIX)
	log.Printf("%d tickets for event with ID %s have been booked", tickets, req.ID)
}
