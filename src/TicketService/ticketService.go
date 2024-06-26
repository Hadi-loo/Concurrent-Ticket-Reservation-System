package TicketService

import (
	"Ticket_Resevation/src/Event"
	"Ticket_Resevation/src/Utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type TicketService struct {
	Events sync.Map
	mu     sync.RWMutex
}

func InitTicketService(path string, fileName string, ticketService *TicketService) {

	ticketService.mu.Lock()
	defer ticketService.mu.Unlock()

	// for Linux
	data, err := os.ReadFile(path + fileName)
	// for Windows
	//data, err := os.ReadFile(fileName)

	if err != nil {
		log.Println("Error reading file: ", err)
		return
	}

	var jsonData struct {
		Events []Event.Event `json:"events"`
	}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		log.Println("Error unmarshalling JSON: ", err)
		return
	}

	for _, event := range jsonData.Events {
		ticketService.Events.Store(event.ID, &event)
	}

	log.Println("Ticket Service initialized")
}

func Save(path string, fileName string, ticketService *TicketService) {

	ticketService.mu.Lock()
	defer ticketService.mu.Unlock()

	var events []Event.Event
	ticketService.Events.Range(func(key, value interface{}) bool {
		event := value.(*Event.Event)
		events = append(events, *event)
		return true
	})

	data, err := json.Marshal(map[string]interface{}{"events": events})
	if err != nil {
		log.Println("Error marshalling JSON: ", err)
		return
	}

	// for Linux
	err = os.WriteFile(path+fileName, data, 0644)
	// for Windows
	//os.WriteFile(fileName, data, 0644)

	if err != nil {
		log.Println("Error writing file: ", err)
		return
	}

	log.Println("Ticket Service saved")
}

func (ts *TicketService) CreateEvent(name string, date time.Time, totalTickets int) (*Event.Event, error) {

	ts.mu.Lock()
	defer ts.mu.Unlock()

	event := &Event.Event{
		ID:               Utils.GenerateUUID(),
		Name:             name,
		Date:             date,
		TotalTickets:     totalTickets,
		AvailableTickets: totalTickets,
		Mu:               sync.RWMutex{},
	}

	ts.Events.Store(event.ID, event)
	return event, nil
}

func (ts *TicketService) ListEvents() []*Event.Event {

	ts.mu.RLock()
	defer ts.mu.RUnlock()

	var events []*Event.Event
	ts.Events.Range(func(key, value interface{}) bool {
		event := value.(*Event.Event)
		events = append(events, event)
		return true
	})
	return events
}

func (ts *TicketService) BookTickets(eventID string, numTickets int) ([]string, error) {

	ts.mu.RLock()
	defer ts.mu.RUnlock()

	event, ok := ts.Events.Load(eventID)
	if !ok {
		return nil, fmt.Errorf("event with ID %s not found", eventID)
	}

	eventObj := event.(*Event.Event)
	eventObj.Mu.Lock()
	defer eventObj.Mu.Unlock()
	if eventObj.AvailableTickets < numTickets {
		return nil, fmt.Errorf("not enough tickets available for event %s", eventID)
	}

	var ticketIDs []string
	for i := 0; i < numTickets; i++ {
		ticketID := Utils.GenerateUUID()
		ticketIDs = append(ticketIDs, ticketID)
		// FIXME: Save ticketID in some data structure

	}

	eventObj.AvailableTickets -= numTickets
	ts.Events.Store(eventID, eventObj)

	return ticketIDs, nil
}
