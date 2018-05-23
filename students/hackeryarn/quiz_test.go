package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/gophercises/quiz/students/hackeryarn/problem"
	"github.com/gophercises/quiz/students/hackeryarn/quiz"
)

func TestReadCSV(t *testing.T) {
	input := "7+3,10\n1+1,2"
	reader := bytes.NewBufferString(input)

	record1 := []string{"7+3", "10"}
	record2 := []string{"1+1", "2"}
	problems := []problem.Problem{
		problem.New(record1),
		problem.New(record2),
	}

	want := quiz.New(problems)
	got := ReadCSV(reader)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected to read in %v got %v", want, got)
	}
}
