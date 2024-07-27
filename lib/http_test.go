package lib

import "testing"

func TestNewRoundRobinServer(t *testing.T) {
	_, err := NewRoundRobinServer("localhost:8080")
	if err != nil {
		t.Errorf("Error creating round robin server: %v", err)
	}
}