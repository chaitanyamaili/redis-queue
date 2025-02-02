package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chaitanyamaili/redis-queue/queue"
)

func main() {
	if os.Getenv("REDIS_HOST") == "" {
		log.Fatal("REDIS_HOST environment variable is not set")
	}
	if os.Getenv("REDIS_PORT") == "" {
		log.Fatal("REDIS_PORT environment variable is not set")
	}
	if os.Getenv("REDIS_QUEUE") == "" {
		log.Fatal("REDIS_QUEUE environment variable is not set")
	}
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")) //"localhost:6379"
	queueName := os.Getenv("REDIS_QUEUE")

	q := queue.NewQueue(redisAddr, queueName)
	defer q.Close()

	// Enqueue messages
	err := q.Enqueue("Task 1")
	if err != nil {
		log.Fatalf("Enqueue error: %v", err)
	}
	err = q.Enqueue("Task 2")
	if err != nil {
		log.Fatalf("Enqueue error: %v", err)
	}

	queuesize, err := q.QueueSize()
	if err != nil {
		log.Fatalf("Queue size error: %v", err)
	}
	fmt.Println("Queue size:", queuesize)

	// Create a loop over the items present in the queue.
	for {
		msg, err := q.Peek()
		if err != nil {
			log.Fatalf("Peek error: %v", err)
		}
		fmt.Println("Peeked message:", msg)
		if msg == "" {
			break
		}
		_, err = q.Dequeue()
		if err != nil {
			log.Fatalf("Dequeue error: %v", err)
		}
	}
}
