package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://www.google.com"

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("Error: %s is offline. Details: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Success! %s is online with status %d\n", url, resp.StatusCode)
	} else {
		fmt.Printf("Warning: %s returned status code %d\n", url, resp.StatusCode)
	}
}
