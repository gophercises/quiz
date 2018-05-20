package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/hackeryarn/quiz/students/hackeryarn/problem"
)

func TestReadCSV(t *testing.T) {
	problems := "7+3,10\n1+1,2"
	readCloser := bytes.NewBufferString(problems)

	record1 := []string{"7+3", "10"}
	record2 := []string{"1+1", "2"}
	want := []problem.Problem{
		problem.New(record1),
		problem.New(record2),
	}
	got := ReadCSV(readCloser)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected to read in %v got %v", want, got)
	}
}
