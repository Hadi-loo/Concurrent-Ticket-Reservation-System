# Concurrent Ticket Reservation System

In this project, we have developed a server-client model for ticket reservation using the Go programming language. 
The server possesses the capability to manage multiple client requests concurrently, thus enabling concurrency control over both server and client connections. 
Additionally, we have incorporated a caching system to optimize the handling of regularly recurring requests.

- [Requirements](#requirements)
- [Overview](#overview)
  - [Concurrent Request Handling](#concurrent-request-handling)
  - [Ticket Reservation Concurrency Control](#ticket-reservation-concurrency-control)
  - [Resource Management](#resource-management)
  - [Fairness and Scheduling](#fairness-and-scheduling)
  - [Caching System](#caching-system)
- [Structure and Implementation](#structure-and-implementation)
- [Results](#results)
  - [Start Terminal](#start-terminal)
  - [Help Command](#help-command)
  - [List Command](#list-command)
  - [Book Command](#book-command)
  - [Create Command](#create-command)
- [How to run](#how-to-run)
  - [1. Compile Dependencies](#1-compile-dependencies)
  - [2. Run Server](#2-run-server)
  - [3. Run Client](#3-run-client)
- [Contributions](#contributions)

## Requirements

For this project, the installation of the Go programming language is required. also we need uuid package inorder to generate uniquely identify entities without centralized coordination:

```go
module Ticket_Resevation

go 1.22

require github.com/google/uuid v1.6.0
```
You can use the following command to download the requirments metioned in box above:

```go
go mod download
```

## Overview
As we mentioned before, the project is a server-client model for ticket reservation. 
The connection between the server and the client is established using the HTTP protocol via the `net/http` package in Go. 

### Concurrent Request Handling

This server is capable of handling multiple client requests concurrently by utilizing goroutines. The net/http package in Go handles concurrent requests to a server through the use of Goroutines. 
When a request is received by an HTTP server, it typically spawns a new Goroutine to handle that request.  
This allows the server to handle multiple requests concurrently without blocking. 

When you call `http.ListenAndServe` to start an HTTP server, it internally calls the Server method on an `http.Server` instance. This method starts listening for incoming HTTP requests on the specified address.  
When an HTTP request is received by the server, the `Server` method spawns a new Goroutine to handle that request. This allows the server to continue accepting and handling new requests without waiting for the current request to complete.  
Each incoming request is handled by invoking the `ServeHTTP` method on an instance of a handler (usually an object that implements the `http.Handler` interface). This method is called in a Goroutine, allowing it to execute concurrently with other requests.  
Go's scheduler efficiently manages a pool of Goroutines, allowing them to run concurrently on available CPU cores. As a result, multiple requests can be processed simultaneously without blocking.  
The response generated by the handler function is written back to the client using the underlying network connection. Go's buffered I/O and non-blocking I/O operations ensure efficient and concurrent writing of responses to multiple clients.  
The `net/http` package includes support for request cancellation and timeouts. This ensures that long-running requests or unresponsive clients do not block the server indefinitely. Goroutines can be cancelled, and timeouts can be set to handle such scenarios gracefully. 

### Ticket Reservation Concurrency Control

In this project, we have to implement a concurrency control mechanism to prevent problems like race conditions, deadlocks, resource contention, and data inconsistency.  
To achieve this, we need to use synchronization primitives like mutexes, channels, atomic operations, and other concurrency control mechanisms provided by Go. Also, we need to identify critical sections of code that need to be protected from concurrent access.

For this purpose, we used mutexes to protect shared resources like the map of events and the number of available tickets.  
```go
type TicketService struct {
	Events sync.Map
	mu     sync.RWMutex
}
```
The `sync.Map` type is a concurrent map that can be safely accessed by multiple Goroutines. It provides atomic operations for reading, writing, and deleting key-value pairs.  
But it's not enough to protect the map of events. Because we are storing the pointer to the event in the map, we need to use a mutex to protect the event itself.  
```go
type Event struct {
	ID               string
	Name             string
	Date             time.Time
	TotalTickets     int
	AvailableTickets int
	Mu               sync.RWMutex
}
```
The `sync.RWMutex` type is a reader-writer mutex that allows multiple readers or a single writer to access a shared resource. It provides exclusive access to the resource for writing and shared access for reading.  
When a Goroutine wants to read the event information, it can acquire a read lock using the `RLock` method. This allows multiple Goroutines to read the event information concurrently.  
When a Goroutine wants to update the event information, it can acquire a write lock using the `Lock` method. This ensures that only one Goroutine can write to the event at a time, preventing data corruption.  
By using mutexes to protect shared resources, we can ensure that concurrent access to the ticket reservation system is safe and sound.

### Resource Management

Because of the lightweight nature of Goroutines, the server can handle a large number of concurrent requests without consuming excessive system resources, so usually we don't need to worry about resource management of Goroutines.  
But in some cases, we may need to limit the number of concurrent Goroutines to prevent resource exhaustion. This can be achieved by using a buffered channel to control the number of active Goroutines. Another approach is to use a semaphore to limit the number of concurrent Goroutines. Also, we can use a worker pool to manage a fixed number of Goroutines that process incoming requests. There are also packages like `netutil` that provide utilities for limiting the number of concurrent Goroutines.

### Fairness and Scheduling

Go provides a fair scheduling mechanism for Goroutines. When multiple Goroutines are competing for CPU time, the Go scheduler ensures that each Goroutine gets a fair share of CPU resources. This prevents any single Goroutine from monopolizing the CPU and starving other Goroutines. Some of the tequniques that Go uses to achieve fair scheduling include:
- **Work Stealing**: The Go scheduler uses a work-stealing algorithm to balance the workload across multiple CPU cores. When a Goroutine is blocked on a channel or mutex, the scheduler can steal work from other Goroutines to keep all CPU cores busy.
- **Preemptive Scheduling**: The Go scheduler uses preemptive scheduling to ensure that no single Goroutine can block the entire program. If a Goroutine is running for too long, the scheduler can preempt it and switch to another Goroutine.
- **Round-Robin Local Run Queue**: Each CPU core has its own local run queue that stores Goroutines waiting to be executed. The scheduler uses a round-robin algorithm to select Goroutines from the local run queue, ensuring that all Goroutines get a fair chance to run.
- **Cooperative Yielding**: Goroutines can voluntarily yield the CPU by calling the `runtime.Gosched` function. This allows other Goroutines to run and prevents any single Goroutine from monopolizing the CPU.
- **Priority-Based Scheduling**: The Go scheduler uses priority-based scheduling to prioritize certain Goroutines over others. For example, the scheduler may give higher priority to Goroutines that are waiting on a channel or mutex.
- **Network Poller**: The Go scheduler uses a network poller to efficiently handle I/O operations. This allows the scheduler to multiplex network connections and avoid blocking on I/O operations.
- **Fair Mutexes**: Go provides fair mutexes that ensure that Goroutines waiting on a mutex are granted access in the order they requested it. This prevents any single Goroutine from starving other Goroutines waiting on the same mutex.

### Caching System

To optimize the handling of high number of requests for different events, we have used a mutex on the map of events which only allows one Goroutine to write to the map at a time, but multiple Goroutines can read from the map concurrently.  
Each event in return has its own mutex to protect the event information. This allows multiple Goroutines to read the event information concurrently, but only one Goroutine can update the event information at a time. 

We also used sync.Map to store the events. The `sync.Map` is a concurrent map implementation in Go that provides better performance and scalability compared to traditional maps with `sync.RWMutex`.  
In Go, maps are not safe for concurrent access without proper synchronization. Using `sync.RWMutex` to synchronize read and write operations can lead to excessive use of atomic operations, causing cache contention and reducing performance.

`sync.Map` is a specialized map type designed for concurrent access. It provides efficient read and write operations while minimizing contention. It is commonly used in scenarios where multiple goroutines need to access and modify a shared map concurrently. It is particularly useful when the map is expected to have a high rate of reads and infrequent writes, which is the case in our ticket reservation system.

Here is the underlying structure of the `sync.Map`:
```go
type Map struct {
   mu     Mutex
   read   atomic.Value // readOnly
   dirty  map[interface{}]*entry
   misses int
}
```
```go
type readOnly struct {
   m       map[interface{}]*entry
   amended bool
}
```

The `sync.Map` type uses a combination of a mutex, an atomic value, and a map to provide concurrent access to the map. The `mu` field is a mutex that protects the map from concurrent access. The `read` field is an atomic value that stores a read-only view of the map. The `dirty` field is a map that stores entries that have been modified. The `misses` field is a counter that tracks the number of cache misses.  
In fact, `sync.Map` maintains two underlying maps: a read-only map (`read`) and a write-only map (`dirty`). The read-only map contains the most recently accessed entries and is optimized for read operations. The write-only map holds entries that have been modified or newly added and are not yet present in the read-only map.

When a goroutine wants to read from the map (by calling methods like `Load` or `Range`), it first checks the `read` field to see if it has a read-only view of the map. If it does, it reads from the read-only view and the corresponding value is returned immediately. If it doesn't, it acquires a lock on the mutex and the write-only map is checked. If the key exists there, the entry is promoted to the read-only map, and the value is returned.

When a goroutine wants to write to the map (by calling methods like `Store` or `Delete`), it acquires a lock on the mutex and writes to the `dirty` map. This ensures that only one goroutine can write to the map at a time. When the dirty map reaches a certain size, the goroutine merges the dirty map with the read-only view and updates the atomic value. This allows the changes to be visible to other goroutines.

The `sync.Map` employs a caching strategy to optimize read operations. When a key is loaded, the corresponding entry is moved from the write-only map to the read-only map, making subsequent reads faster. This caching mechanism ensures that frequently accessed entries are readily available in the read-only map, reducing contention and improving performance.

## Structure and Implementation

A Concurrent Ticket Reservation System is designed to handle multiple users simultaneously attempting to book tickets or reserve seats for events. These systems need to manage seat availability, prevent conflicts, and ensure smooth user experiences. Here, we explain the structure and implementation of our simulated ticket reservation system:

### Ticket

We have a simple Go struct named `Ticket`:

```go
type Ticket struct {
	ID      string
	EventID string
}
```

It has two fields:

- `ID`: A string representing the unique identifier for the ticket.
- `EventID`: A string representing the unique identifier for the event associated with this ticket.

### Event

The `Event` struct represents an event. It is likely part of a larger system (ticket reservation system) where events need to be tracked.

```go
type Event struct {
	ID               string
	Name             string
	Date             time.Time
	TotalTickets     int
	AvailableTickets int
	Mu               sync.RWMutex
}
```

`Event` has several fields:

- `ID`: A string representing the unique identifier for the event.
- `Name`: A string containing the name or title of the event.
- `Date`: A `time.Time` value indicating the date and time of the event.
- `TotalTickets`: An integer representing the total number of tickets for the event.
- `AvailableTickets`: An integer representing the number of tickets currently available (not yet booked or reserved).
- `Mu`: A sync.RWMutex (read-write mutex) used for managing concurrent access to the event data.

The `Mu` field is a synchronization primitive that allows multiple readers or a single writer to access the event data concurrently. It ensures thread safety when reading or modifying the event properties. For example, when booking a ticket, the system would acquire a write lock (`Mu.Lock()`) to update the AvailableTickets count atomically. Similarly, when checking available tickets, the system would acquire a read lock (`Mu.RLock()`) to prevent concurrent writes during the read operation.

### Database

Our database consists of a JSON file named `events.json`. Here is an example of the contents of this file:

```json
{
  "events": [
    {
      "ID": "1",
      "Name": "Event 1",
      "Date": "2002-03-11T00:00:00Z",
      "TotalTickets": 100,
      "AvailableTickets": 50
    },
    {
      "ID": "2",
      "Name": "Event 2",
      "Date": "2002-03-12T00:00:00Z",
      "TotalTickets": 200,
      "AvailableTickets": 100
    },
  ]
}
```

The JSON file describes three events, each with an ID, name, date, total ticket count, and available ticket count. It’s a simple representation of event data.

### TicketService

The `TicketService` struct represents a service responsible for managing ticket-related operations. It has two fields:

- `Events`: A sync.Map used to store event data (keyed by event ID).
- `mu`: A `sync.RWMutex` (read-write mutex) used for managing concurrent access to the Events map.

Now we will examine the functions and methods related to this service:

`InitTicketService()`:

This method initializes the `TicketService` by reading event data from a JSON file.

- Acquires a write lock (`mu.Lock()`) to ensure exclusive access during initialization.
- Converting the data into a struct that contains an array of events.
- Stores each event in the Events map using its ID as the key.
- Logs a message indicating successful initialization.
  
`Save()`:

The `Save` method saves the current state of the `TicketService` (events) back to a JSON file.

- Iterates through the events stored in the `Events` map.
- Acquires a write lock (`mu.Lock()`) to prevent concurrent modifications during saving.
- Writes the JSON data to the specified file (either for Linux or Windows).
- Logs a message indicating successful saving.

`CreateEvent()`:

This method creates a new event and adds it to the `Events` map.

- Creates a new event, initializes the event with the provided name, date, and total ticket count, and adds it to the `Events` map.
- Acquires a write lock (`ts.mu.Lock()`) to ensure exclusive access during event creation.
- Generates a unique event ID using a utility function (`Utils.GenerateUUID()`).
- Returns the created event and a potential error.

`ListEvents()`:

This function retrieves a list of all events stored in the Events map.

- Acquires a read lock (`ts.mu.RLock()`) to allow concurrent reading.
- Initializes an empty slice of `Event` pointers (events).
- Iterates through the events stored in the `Events` map using `ts.Events.Range` and appends each event to the events slice.
- Returns the list of events.

`BookTickets()`:

The `BookTickets()` function allows booking a specified number of tickets for a given event.

- Acquires a read lock (`ts.mu.RLock()`) to allow concurrent reading.
- Loads the event associated with the given eventID from the `Events` map.
- Acquires a write lock on the event (`eventObj.Mu.Lock()`) to prevent concurrent modifications.
- Checks if there are enough available tickets for booking and generates unique ticket IDs for the booked tickets.
- Decrements the available ticket count for the event and stores the updated event back in the `Events` map.
- Returns the list of booked ticket IDs and a potential error.

### Server

Here, we describe the specifications and features of our HTTP server. This server handles HTTP requests from clients.

Let’s explore the functions and methods associated with this server:

`startServer()`:

It starts the HTTP server to listen on port 8080.

- Sets the port variable to “:8080” and Configures the log prefix to use the value of `LOGS_INFO_PREFIX`.
- Calls `http.ListenAndServe(port, nil)` to start the server.
- Logs a message indicating that the server has started on the specified port.
- If any error occurs during server startup, it logs the error and terminates the program with `log.Fatal`.

`ListEventsHandler()`:

the `ListEventsHandler()` function handles HTTP requests for listing events. In this context:

`w` refers to an `http.ResponseWriter` used to write the HTTP response.
`r` an `http.Request` representing the incoming HTTP request.
`ticketService` is a pointer to the `TicketService` instance.

**Purpose**: Retrieves a list of events from the server.
**Request and Response Body Format**: JSON

- Calls the `ListEvents` method of the `ticketService` to retrieve a list of events.
- Acquires read locks (`events[i].Mu.RLock()`) for each event.
- Sets the HTTP response header to indicate that the content type is JSON.
- Writes the JSON response to the `http.ResponseWriter`.
- Logs a message indicating that the `/list-events` endpoint has been called.

`CreateEventHandler()`:

It handles HTTP requests for creating a new event.

**Purpose**: Creating an event, and returning the event details in JSON format.
**Request and Response Body Format**: JSON

- Decodes the incoming JSON request body into an instance of `CreateEventRequest`.
- Calls the `CreateEvent` method of the `ticketService` to create a new event.
- Sets the HTTP response header to indicate that the content type is JSON.
- Writes the JSON response to the `http.ResponseWriter`.
- Logs a message indicating the successful creation of the new event.

`BookTicketsHandler()`:

This function handles HTTP requests for booking tickets.

**Purpose**: Booking the specified number of tickets, and returning the booked ticket IDs in JSON format.
**Request and Response Body Format**: JSON

- Decodes the incoming JSON request body into an instance of `BookTicketsRequest`.
- Calls the `BookTickets` method of the `ticketService` to book the specified number of tickets for the given event.
- Sets the HTTP response header to indicate that the content type is JSON.
- Writes the JSON response to the `http.ResponseWriter`.
- Logs a message indicating the successful booking of tickets.


## Results

The outcomes of each code functionality are displayed in the images below:

### Start Terminal

#### Server
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/5ebfdd4c-9b02-4959-badf-674d4ca7e14b)


#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/e262f8b1-97d1-4d9e-aaae-4a8787f841d9)



### Help Command


With the following command, you have the capability to view the syntax and structure of each individual command:
```txt
help
```

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/158e6763-2eb6-4706-b78e-7c9ae0e72d0f)


### List Command

With the following command, you can observe all Events and their information:

```txt
list
```

#### Client 
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/113b8344-80de-4e83-a70b-a9446aa6a909)


#### Server 
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/e52ee5c7-2b48-49b2-be3e-16713d07494b)


### Book Command

With the following command, you can book tickets for your intended Event:

```txt
book [Event_ID] [Number of Tickets]
```

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/0589991f-554b-4490-8bf9-412452f6eae5)


![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/bd40e052-15cf-462a-925e-03b99b512b8b)



#### server
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/172ffb1a-bc29-4aa9-8137-3ea62235a17b)


### Create Command

```txt
create [Name] [Date] [Total Tickets]
```

With the following command, you can generate an Event with the specified details:

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/fd9791f6-549d-4a05-b628-57648458def6)

![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/3257cf71-a286-4e5d-9e33-e81e02b8bc0f)



#### Server
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/c6dc7f72-8ce1-4fe0-b341-d4852b4ff878)



## How to run

To run this project, follow these steps:

### 1. Compile Dependencies

Ensure that you have Go installed on your system. Navigate to the root directory of the project where the `go.mod` file is located, and compile the dependencies by executing:

```bash
go mod download
```

### 2. Run Server

To start the server, run the following command:

```bash
go run server/server.go
```

The server will start listening for incoming clients request.

### 3. Run Client

To execute the client code, run the following command:

```bash
go run client/client.go
```

This will send requests to the server and display the responses.


## Contributions

- **Server File**: Hadi Babaloo, Sina Tabasi
- **Client File**: Hadi Babaloo, Kasra Haji-Heydari
- **proto File**: Hadi Babaloo, MohammadSadegh Aboofazeli
- **Report and Analysis**: MohammadSadegh Aboofazeli, Sina Tabasi, Kasra Haji-Heydari
