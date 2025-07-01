/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2025/6/27 -- 14:33
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xqueue xqueue/queue.go
*/

package xqueue

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrQueueClosed = errors.New("queue is closed")
)

// Message represents a message in the queue.
type Message struct {
	// ID is the unique identifier of the message.
	ID string
	// Body is the content of the message.
	Body any
}

// Queue is a thread-safe FIFO queue.
type Queue struct {
	lock     sync.Mutex
	cond     *sync.Cond
	messages []*Message
	closed   bool
}

// NewQueue creates and returns a new Queue.
func NewQueue() *Queue {
	q := &Queue{}
	q.cond = sync.NewCond(&q.lock)
	return q
}

// Push adds a message to the back of the queue.
// It returns ErrQueueClosed if the queue is closed.
func (q *Queue) Push(body any) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.closed {
		return ErrQueueClosed
	}

	msg := &Message{
		ID:   uuid.NewString(),
		Body: body,
	}
	q.messages = append(q.messages, msg)

	// Signal a waiting consumer that a new message is available.
	q.cond.Signal()

	return nil
}

// Pop removes and returns a message from the front of the queue.
// It blocks if the queue is empty until a message is available or the queue is closed.
// It returns ErrQueueClosed if the queue is empty and closed.
func (q *Queue) Pop() (*Message, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for len(q.messages) == 0 {
		if q.closed {
			return nil, ErrQueueClosed
		}
		// Wait for a new message to be pushed.
		q.cond.Wait()
	}

	msg := q.messages[0]
	q.messages = q.messages[1:]

	return msg, nil
}

// TryPop removes and returns a message from the front of the queue if available.
// It returns immediately with a boolean indicating if a message was retrieved.
func (q *Queue) TryPop() (*Message, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.messages) == 0 || q.closed {
		return nil, false
	}

	msg := q.messages[0]
	q.messages = q.messages[1:]

	return msg, true
}

// Len returns the current number of messages in the queue.
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.messages)
}

// Close closes the queue and wakes up all waiting consumers.
// After closing, no more messages can be pushed to the queue.
func (q *Queue) Close() {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.closed {
		return
	}

	q.closed = true
	// Broadcast to all waiting consumers to wake them up.
	q.cond.Broadcast()
}
