package lib

import "testing"

func TestReadConfig(t *testing.T) {
	wants := []string{"abc", "def", "hij"}

	configReader := NewConfigReader("servers.json")

	got, err := configReader.ReadConfig()
	if err != nil {
		t.Error("Failed to read config")
	}

	for i, want := range wants {
		if want != got[i] {
			t.Errorf("Got %s, want %s", got[i], want)
		}
	}
}