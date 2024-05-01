# Concurrent Ticket Reservation System

- [Requirements](#requirements)
- [Structure and Implementation](#structure-and-implementation)
- [Results](#results)
- [How to run](#how-to-run)
- [Contributions](#contributions)

## Requirements

For this project, the installation of the Go programming language is required. Additionally, the dependencies for gRPC packages are specified in the go.mod file, as illustrated in the following box:

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

To obtain results for all four methods, the server and client files are executed, after which requests are dispatched for each method accordingly:

### Start Terminal

#### Server
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/df2fbd1c-e301-41a2-928d-78eda1d695a7)

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/24d28151-7175-4732-9e10-8806cea81905)


### Help Command


With the following command, you have the capability to view the syntax and structure of each individual command:
```txt
help
```

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/d6f740dc-8495-44e3-a098-6b04cb5f55c3)

### List Command

With the following command, you can observe all Events and their information:

```txt
list
```

#### Client 
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/ac9dd503-e13b-4ea5-989b-03bf34acdd0a)

#### Server 
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/b52c6f08-3f76-49ab-b586-4949ae822d40)

### Book command

With the following command, you can book tickets for your intended Event:

```txt
book [Event_ID] [Number of Tickets]
```

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/46a4b448-819d-48b2-a83a-f1c26fabf94a)

![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/a66837c8-f7ad-42b0-ae77-258e25a60682)


#### server
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/f15ffbd2-6806-4a59-904c-ca892d982f29)

### Create command

With the following command, you can generate an Event with the specified details:

#### Client
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/86a3d389-79e0-4fcd-a344-65d829eed52e)
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/b74e7c38-6b8d-4344-b5c1-46faa40350bb)


#### Server
![image](https://github.com/Hadi-loo/Concurrent-Ticket-Reservation-System/assets/88041997/4e19cc55-755b-41f9-9d3a-92ec98c36e14)


## How to run

To run this project, follow these steps:

### 1. Compile Dependencies

Ensure that you have Go installed on your system. Navigate to the root directory of the project where the `go.mod` file is located, and compile the dependencies by executing:

```bash
go mod download
```

### 3. Run Server

To start the server, run the following command:

```bash
go run server/server.go
```

The server will start listening for incoming clients request.

### 4. Run Client

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
