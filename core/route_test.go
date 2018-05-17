package core

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	t.Run("Returns200WhenStatusNotSet", func(t *testing.T) {
		expectedStatus := 200
		route := Route{}
		result := route.GetStatus()
		if result != expectedStatus {
			t.Fatalf("Expected %d, got %d", expectedStatus, result)
		}
	})

	t.Run("ReturnsStatusWhenSet", func(t *testing.T) {
		expectedStatus := 400
		route := Route{Status: 400}
		result := route.GetStatus()
		if result != expectedStatus {
			t.Fatalf("Expected %d, got %d", expectedStatus, result)
		}
	})
}
