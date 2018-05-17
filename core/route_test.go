package core

import (
	"testing"
)

func TestGetStatus_Returns200WhenStatusNotSet(t *testing.T) {
	expectedStatus := 200
	route := Route{}
	result := route.GetStatus()
	if result != expectedStatus {
		t.Fatalf("Expected %d, got %d", expectedStatus, result)
	}
}

func TestGetStatus_ReturnsStatusWhenSet(t *testing.T) {
	expectedStatus := 400
	route := Route{Status: 400}
	result := route.GetStatus()
	if result != expectedStatus {
		t.Fatalf("Expected %d, got %d", expectedStatus, result)
	}
}
