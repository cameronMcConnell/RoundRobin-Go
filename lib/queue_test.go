package lib

import "testing"

func TestNewQueue(t *testing.T) {
	_, err := NewQueue()
	if err != nil {
		t.Errorf("Error creating new queue: %v", err)
	}
}

func TestEnque(t *testing.T) {
	queue, err := NewQueue()
	if err != nil {
		t.Errorf("Error creating new queue: %v", err)
	}

	wants := []string{"abc", "def", "hij", "klm"}

	queue.Enque("klm")

	for _, want := range wants {
		got, err := queue.Deque()
		if err != nil || got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	}
}

func TestDeque(t *testing.T) {
	queue, err := NewQueue()
	if err != nil {
		t.Error(err)
	}

	wants := []string{"abc", "def", "hij"}

	for _, want := range wants {
		got, err := queue.Deque()
		if err != nil || got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	}
}