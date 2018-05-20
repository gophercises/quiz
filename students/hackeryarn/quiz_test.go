package main

import (
	"bytes"
	"testing"
)

func TestReadCSV(t *testing.T) {
	problems := "7+3,10\n1+1,2"
	readCloser := bytes.NewBufferString(problems)

	want := []Problem{
		{"7+3", "10"},
		{"1+1", "2"},
	}
	got := ReadCSV(readCloser)

	if !compareProblems(want, got) {
		t.Errorf("Expected to read in %v got %v", want, got)
	}
}

func compareProblems(a, b []Problem) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
