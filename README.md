# Concurrent Ticket Reservation System

In this project, we have developed a server-client model for ticket reservation using the Go programming language. The server possesses the capability to manage multiple client requests concurrently, thus enabling concurrency control over both server and client connections. Additionally, we have incorporated a caching system to optimize the handling of regularly recurring requests.

- [Requirements](#requirements)
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

## Structure and Implementation
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
