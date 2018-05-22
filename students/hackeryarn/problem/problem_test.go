package problem

import (
	"bufio"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	record := []string{"question", "answer"}

	want := Problem{"question", "answer"}
	got := New(record)

	if got != want {
		t.Errorf("expected to create problem %v got %v", want, got)
	}
}

func TestCheckAnswer(t *testing.T) {
	record := []string{"7+3", "10"}
	problem := New(record)

	t.Run("checks the correct answer", func(t *testing.T) {
		input := strings.NewReader("10\n")

		want := true
		got := problem.CheckAnswer(bufio.NewScanner(input))

		if want != got {
			t.Errorf("Expected to return %v got %v", want, got)
		}
	})

	t.Run("checks incorrect answer", func(t *testing.T) {
		input := strings.NewReader("2\n")

		want := false
		got := problem.CheckAnswer(bufio.NewScanner(input))

		if want != got {
			t.Errorf("Expected to return %v got %v", want, got)
		}
	})
}
