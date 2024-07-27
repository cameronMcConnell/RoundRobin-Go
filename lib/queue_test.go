package lib

import "testing"

func TestEnque(t *testing.T) {
	queue := newQueue()

	vals := []string{"hello", "world", "noo"}

	for _, val := range vals {
		queue.Enque(val)
	}
}

func TestDeque(t *testing.T) {
	queue := newQueue()

	vals := []string{"hello", "world", "yess"}

	for _, val := range vals {
		queue.Enque(val)
	}

	for _, want := range vals {
		got, err := queue.Deque()
		if err != nil || got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	}
}