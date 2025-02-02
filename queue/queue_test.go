package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const redisAddr = "localhost:6379"
const queueName = "testQueue"

func TestQueue(t *testing.T) {
	q := NewQueue(redisAddr, queueName)
	defer q.Close()

	// Test enqueue
	err := q.Enqueue("hello world")
	assert.NoError(t, err)

	// Test queue size
	size, err := q.QueueSize()
	assert.NoError(t, err)
	assert.Equal(t, int64(1), size)

	// Test peek
	msg, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", msg)

	// Test dequeue
	deqMsg, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, "hello world", deqMsg)

	// Test empty queue
	_, err = q.Dequeue()
	assert.Error(t, err)
}
