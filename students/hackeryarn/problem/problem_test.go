package problem

import (
	"testing"
)

func TestNew(t *testing.T) {
	record := []string{"question", "answer"}

	want := Problem{"question", "answer"}
	got := New(record)

	if got != want {
		expected("expected to create problem %v got %v", want, got)
	}
}
