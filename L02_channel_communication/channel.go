package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup
	count := 3

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Worker %d: Sleeping...\n", id)
			time.Sleep(1 * time.Second)

			messages <- fmt.Sprintf("Message from Go routine %d!", id)
			fmt.Printf("Worker %d: Message sent\n", id)
		}(i)
	}

	go func() {
		wg.Wait()
		close(messages)
	}()

	fmt.Println("Main: Waiting for messages...")
	for msg := range messages {
		fmt.Printf("Main: Received %s\n", msg)
	}
}
