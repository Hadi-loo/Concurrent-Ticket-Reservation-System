package Event

import (
	"sync"
	"time"
)

type Event struct {
	ID               string
	Name             string
	Date             time.Time
	TotalTickets     int
	AvailableTickets int
	Mu               sync.RWMutex
}
