package app

import (
	"os"
	"testing"
)

func TestNewApp(t *testing.T) {
	defer os.Remove("todos.txt") //deletes file created while testing
	modes := []string{"sql", "file", "mongo", ""}
	for _, m := range modes {
		if _, err := New(m); err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}
	}
}

func TestNewAppUnsupported(t *testing.T) {
	dataMode := "unsupported"
	e := "Unexpected error: not supported db mode: " + dataMode
	if _, err := New(dataMode); err == nil {
		t.Errorf("Expected error: %v, got %v", e, nil)
		return
	}
}
