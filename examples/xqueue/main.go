package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/xneogo/extensions/xqueue"
)

func main() {
	// Create a new queue
	q := xqueue.NewQueue()

	var wg sync.WaitGroup

	// Start a producer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("message-%d", i)
			if err := q.Push(msg); err != nil {
				log.Printf("Failed to push message: %v", err)
				return
			}
			log.Printf("=> Pushed: %s", msg)
			time.Sleep(100 * time.Millisecond)
		}
		// After finishing pushing, close the queue to signal consumers
		log.Println("Producer finished, closing queue.")
		q.Close()
	}()

	// Start consumer goroutines
	numConsumers := 3
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func(consumerID int) {
			defer wg.Done()
			for {
				msg, err := q.Pop()
				if err != nil {
					if err == xqueue.ErrQueueClosed {
						log.Printf("Consumer %d: Queue is closed, exiting.", consumerID)
						return
					}
					log.Printf("Consumer %d: Failed to pop message: %v", consumerID, err)
					return
				}
				log.Printf("<= Consumer %d received: %s (ID: %s)", consumerID, msg.Body, msg.ID)
				time.Sleep(300 * time.Millisecond) // Simulate work
			}
		}(i + 1)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("All goroutines finished. Exiting.")
}
