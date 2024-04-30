package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func SetupClient(client *http.Client) {
	client.Timeout = 10 * time.Second
}

func ListEventsCommandHandler(client *http.Client, args []string) []byte {
	if len(args) > 0 {
		log.SetPrefix(LOGS_WARN_PREFIX)
		log.Println("Unknown arguments for list command")
	}
	req, err := http.NewRequest("GET", ListEventsUrl, nil)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error creating request: " + err.Error())
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error sending request: " + err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error reading response: " + err.Error())
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
		log.Printf(string(body))
		return nil
	}

	return body
}

func CreateEventCommandHandler(client *http.Client, args []string) []byte {
	if len(args) != 3 {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Invalid arguments for create command")
		return nil
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"name":         args[0],
		"date":         args[1],
		"totalTickets": args[2],
	})
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error marshalling request: " + err.Error())
		return nil
	}

	req, err := http.NewRequest("POST", CreateEventUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error creating request: " + err.Error())
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error sending request: " + err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error reading response: " + err.Error())
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
		log.Printf(string(body))
		return nil
	}

	return body
}

func BookTicketsCommandHandler(client *http.Client, args []string) []byte {
	if len(args) != 2 {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Invalid arguments for book command")
		return nil
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"eventID":    args[0],
		"numTickets": args[1],
	})
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error marshalling request: " + err.Error())
		return nil
	}

	req, err := http.NewRequest("POST", BookTicketsUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error creating request: " + err.Error())
		return nil
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error sending request: " + err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Println("Error reading response: " + err.Error())
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.SetPrefix(LOGS_ERROR_PREFIX)
		log.Printf("Unexpected status code: %d\n", resp.StatusCode)
		log.Print(string(body))
		return nil
	}

	return body
}
