package core

import "testing"

func TestMatching(t *testing.T) {
	t.Run("MatchPath_works_as_expected", func(t *testing.T) {
		result := MatchPath()

		if result != expectedStatus {
			t.Fatalf("Expected %d, got %d", expectedStatus, result)
		}
	})
}
