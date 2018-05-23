package main

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/gophercises/quiz/students/hackeryarn/problem"
	"github.com/gophercises/quiz/students/hackeryarn/quiz"
)

type flaggerMock struct {
	stringVarCalls  int
	stringVarNames  []string
	stringVarValues []string
	stringVarUsages []string
}

func (f *flaggerMock) StringVar(p *string, name, value, usage string) {
	f.stringVarCalls++
	f.stringVarNames = append(f.stringVarNames, name)
	f.stringVarValues = append(f.stringVarValues, value)
	f.stringVarUsages = append(f.stringVarUsages, usage)
}

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
		t.Errorf("it should read in %v got %v", want, got)
	}
}

func TestConfigFlags(t *testing.T) {
	flagger := &flaggerMock{}

	ConfigFlags(flagger)

	assertStringVars(t, flagger)
}

func assertStringVars(t *testing.T, flagger *flaggerMock) {
	t.Helper()

	if flagger.stringVarCalls != 1 {
		t.Errorf("it should call StringVars %d times, called %d", 1, flagger.stringVarCalls)
	}

	expectedNames := []string{FileFlag}
	expectedValues := []string{FileFlagValue}
	expectedUsages := []string{FileFlagUsage}

	if !reflect.DeepEqual(expectedNames, flagger.stringVarNames) {
		t.Errorf("it should setup StringVar names to be %v, got %v", expectedNames, flagger.stringVarNames)
	}

	if !reflect.DeepEqual(expectedValues, flagger.stringVarValues) {
		t.Errorf("it should setup StringVar values to be %v, got %v", expectedValues, flagger.stringVarValues)
	}

	if !reflect.DeepEqual(expectedUsages, flagger.stringVarUsages) {
		t.Errorf("it should setup StringVar usages to be %v, got %v", expectedUsages, flagger.stringVarUsages)
	}
}
