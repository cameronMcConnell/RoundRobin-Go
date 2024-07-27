package lib

import "errors"

type QueueNode struct {
	Value string
	next *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func NewQueue() (*Queue, error) {
	configReader := NewConfigReader("servers.json")

	servers, err := configReader.ReadConfig()
	if err != nil {
		return nil, err
	}

	queue := &Queue{}

	for _, server := range servers {
		queue.Enque(server)
	}

	return queue, nil
}

func (q *Queue) Enque(value string) {
	node := &QueueNode{value, nil}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

func (q *Queue) Deque() (string, error) {
	if q.head == nil {
		return "", errors.New("Queue is empty")
	}

	temp := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	temp.next = nil

	return temp.Value, nil
}