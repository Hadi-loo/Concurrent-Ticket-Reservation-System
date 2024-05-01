package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ExitCommandHandler(args []string) {
	if len(args) > 0 {
		log.SetPrefix(LOGS_WARN_PREFIX)
		log.Println("Unknown arguments for exit command")
	}
	fmt.Println("Exiting...")
	os.Exit(0)
}

func HelpCommandHandler(args []string) {
	if len(args) == 0 {
		fmt.Println("Available commands:")
		fmt.Println("help [command] - Display help for a command")
		fmt.Println("list - List all events")
		fmt.Println("create [name] [date] [totalTickets] - Create a new event")
		fmt.Println("book [id] [tickets] - Book tickets for an event")
		fmt.Println("auto - Run a command automatically multiple times")
		fmt.Println("exit - Exit the program")
		return
	} else if len(args) > 1 {
		log.SetPrefix(LOGS_WARN_PREFIX)
		log.Println("Unknown arguments for help command")
	}
	switch args[0] {
	case "help":
		fmt.Println("help [command] - Display help for a command")
		fmt.Println("Example: help list")
	case "list":
		fmt.Println("list - List all events")
		fmt.Println("Example: list")
	case "create":
		fmt.Println("create [name] [date] [totalTickets] - Create a new event")
		fmt.Println("name - Name of the event")
		fmt.Println("date - Date of the event in RFC3339 format")
		fmt.Println("totalTickets - Total number of tickets available for the event")
		fmt.Println("Example: create EventName 2022-01-01T12:00:00Z 100")
	case "book":
		fmt.Println("book [id] [tickets] - Book tickets for an event")
		fmt.Println("id - ID of the event")
		fmt.Println("tickets - Number of tickets to book")
		fmt.Println("Example: book 1 5")
	case "auto":
		fmt.Println("auto list [requests] [delay] - List events automatically")
		fmt.Println("auto book [id] [tickets] [requests] [delay] - Book tickets automatically")
		fmt.Println("requests - Number of requests to make")
		fmt.Println("delay - Delay between requests in milliseconds")
		fmt.Println("id - ID of the event")
		fmt.Println("tickets - Number of tickets to book")
		fmt.Println("Example: auto list 5 1000")
	case "exit":
		fmt.Println("exit - Exit the program")
		fmt.Println("Example: exit")
	default:
		fmt.Println("Unknown command: " + args[0])
	}
}

func AutoCommandHandler(client *http.Client, args []string) {
	if len(args) < 1 {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Invalid arguments for auto command")
		return
	}
	if args[0] == "list" {
		// the second argument is number of requests and the third is the delay
		if len(args) != 3 {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid arguments for auto list command")
			return
		}
		numRequests, err := strconv.Atoi(args[1])
		if err != nil {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid number of requests: " + err.Error())
			return
		}
		delay, err := strconv.Atoi(args[2])
		if err != nil {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid delay: " + err.Error())
			return
		}
		for i := 0; i < numRequests; i++ {
			res := ListEventsCommandHandler(client, nil)
			PrintListCommandResponse(res)
			if i < numRequests-1 {
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
		}
	} else if args[0] == "book" {
		// the second argument is event ID, the third is number of tickets, the fourth is number of requests and the fifth is the delay
		if len(args) != 5 {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid arguments for auto book command")
			return
		}
		eventID := args[1]
		numTickets, err := strconv.Atoi(args[2])
		if err != nil {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid number of tickets: " + err.Error())
			return
		}
		numRequests, err := strconv.Atoi(args[3])
		if err != nil {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid number of requests: " + err.Error())
			return
		}
		delay, err := strconv.Atoi(args[4])
		if err != nil {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Invalid delay: " + err.Error())
			return
		}
		for i := 0; i < numRequests; i++ {
			res := BookTicketsCommandHandler(client, []string{eventID, strconv.Itoa(numTickets)})
			PrintBookCommandResponse(res)
			if i < numRequests-1 {
				time.Sleep(time.Duration(delay) * time.Millisecond)
			}
		}
	}
}

func UnknownCommandHandler(command string) {
	log.SetPrefix(LOGS_ERROR_PREFIX)
	log.Println("Unknown command: " + command)
}

func PrintListCommandResponse(body []byte) {
	if body == nil {
		return
	}

	var events []map[string]interface{}
	err := json.Unmarshal(body, &events)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error unmarshalling response: " + err.Error())
		return
	}

	for _, event := range events {
		// print in magenta
		fmt.Println("\033[35m------------------------------------------------\033[0m")
		fmt.Println("ID:", event["ID"])
		fmt.Println("Name:", event["Name"])
		fmt.Println("Date:", event["Date"])
		fmt.Println("Total Tickets:", event["TotalTickets"])
		fmt.Println("Available Tickets:", event["AvailableTickets"])
	}
	fmt.Println("\033[35m------------------------------------------------\033[0m")
}

func PrintCreateCommandResponse(body []byte) {
	if body == nil {
		return
	}

	var event map[string]interface{}
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error unmarshalling response: " + err.Error())
		return
	}

	fmt.Println("\033[35m------------------------------------------------\033[0m")
	fmt.Println("ID:", event["ID"])
	fmt.Println("Name:", event["Name"])
	fmt.Println("Date:", event["Date"])
	fmt.Println("Total Tickets:", event["TotalTickets"])
	fmt.Println("Available Tickets:", event["AvailableTickets"])
	fmt.Println("\033[35m------------------------------------------------\033[0m")
}

func PrintBookCommandResponse(body []byte) {
	if body == nil {
		return
	}

	var ticketIDs []string
	err := json.Unmarshal(body, &ticketIDs)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error unmarshalling response: " + err.Error())
		return
	}

	fmt.Println("\033[35m------------------------------------------------\033[0m")
	fmt.Println("\033[36mTicket IDs:\033[0m")
	for _, ticketID := range ticketIDs {
		fmt.Println(ticketID)
	}
	fmt.Println("\033[35m------------------------------------------------\033[0m")
}

func GetInput(client *http.Client) {
	reader := bufio.NewReader(os.Stdin)
	log.SetPrefix(LOGS_INFO_PREFIX)
	log.Println("Enter a command or type 'help' for a list of commands")
	for {
		fmt.Print("\033[36m>> \033[0m")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.SetPrefix(LOGS_ERROR_PREFIX)
			log.Println("Error reading input: " + err.Error())
			continue
		}

		text = strings.TrimSpace(text)
		parts := strings.Split(text, " ")
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		}

		switch command {
		case "exit":
			ExitCommandHandler(args)
		case "help":
			HelpCommandHandler(args)
		case "list":
			res := ListEventsCommandHandler(client, args)
			PrintListCommandResponse(res)
		case "create":
			res := CreateEventCommandHandler(client, args)
			PrintCreateCommandResponse(res)
		case "book":
			res := BookTicketsCommandHandler(client, args)
			PrintBookCommandResponse(res)
		case "auto":
			AutoCommandHandler(client, args)
		default:
			UnknownCommandHandler(command)
		}
	}
}
