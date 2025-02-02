package queue

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Queue struct with Redis client
type Queue struct {
	client *redis.Client
	name   string
	ctx    context.Context
}

// NewQueue initializes a Redis queue
func NewQueue(redisAddr, queueName string) *Queue {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	ctx := context.Background()
	return &Queue{
		client: rdb,
		name:   queueName,
		ctx:    ctx,
	}
}

// Enqueue adds a message to the Redis queue
func (q *Queue) Enqueue(message string) error {
	err := q.client.LPush(q.ctx, q.name, message).Err()
	if err != nil {
		return fmt.Errorf("failed to enqueue message: %v", err)
	}
	return nil
}

// Dequeue retrieves and removes a message from the Redis queue
func (q *Queue) Dequeue() (string, error) {
	msg, err := q.client.RPop(q.ctx, q.name).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("queue is empty")
	} else if err != nil {
		return "", fmt.Errorf("failed to dequeue message: %v", err)
	}
	return msg, nil
}

// Peek retrieves a message without removing it
func (q *Queue) Peek() (string, error) {
	msg, err := q.client.LIndex(q.ctx, q.name, -1).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("queue is empty")
	} else if err != nil {
		return "", fmt.Errorf("failed to peek message: %v", err)
	}
	return msg, nil
}

// QueueSize returns the number of messages in the queue
func (q *Queue) QueueSize() (int64, error) {
	size, err := q.client.LLen(q.ctx, q.name).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get queue size: %v", err)
	}
	return size, nil
}

// Close closes the Redis client connection
func (q *Queue) Close() error {
	return q.client.Close()
}
