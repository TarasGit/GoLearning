package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messages := make(chan string)
	var wg sync.WaitGroup // Hilft uns zu warten, bis alle Worker fertig sind
	count := 3            // Wir testen es mal mit 3 Workern

	fmt.Println("Start")

	for i := 1; i <= count; i++ {
		wg.Add(1)         // Ein Worker kommt hinzu
		go func(id int) { // id als Parameter übergeben!
			defer wg.Done() // Signalisiert am Ende: "Ich bin fertig"

			fmt.Printf("Worker %d: Sleeping...\n", id)
			time.Sleep(2 * time.Second)

			messages <- fmt.Sprintf("Message from Go routine %d!", id)
			fmt.Printf("Worker %d: Message sent\n", id)
		}(i)
	}

	// WICHTIG: In einer extra Go-Routine auf die Worker warten
	// und DANN den Channel schließen.
	go func() {
		wg.Wait()
		close(messages)
	}()

	fmt.Println("Main: Waiting for messages...")
	for msg := range messages {
		fmt.Printf("Main: Received %s\n", msg)
	}

	fmt.Println("Main: Alle Nachrichten empfangen, Ende.")
}
