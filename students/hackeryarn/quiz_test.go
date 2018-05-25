package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	quiz "github.com/gophercises/quiz/students/hackeryarn/myquiz"
	"github.com/gophercises/quiz/students/hackeryarn/problem"
)

type flaggerMock struct {
	stringVarCalls  int
	intVarCalls     int
	varNames        []string
	varUsages       []string
	varStringValues []string
	varIntValues    []int
}

func (f *flaggerMock) StringVar(p *string, name, value, usage string) {
	f.stringVarCalls++
	f.varNames = append(f.varNames, name)
	f.varStringValues = append(f.varStringValues, value)
	f.varUsages = append(f.varUsages, usage)
}

func (f *flaggerMock) IntVar(p *int, name string, value int, usage string) {
	f.intVarCalls++
	f.varNames = append(f.varNames, name)
	f.varIntValues = append(f.varIntValues, value)
	f.varUsages = append(f.varUsages, usage)
}

type timerMock struct {
	duration int
}

func (t *timerMock) NewTimer(d time.Duration) *time.Timer {
	t.duration = int(d.Seconds())
	return time.NewTimer(1 * time.Millisecond)
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

	assertStringCalls(t, flagger)
	assertIntCalls(t, flagger)
	assertFlags(t, flagger)
}

func TestStartTimer(t *testing.T) {
	timer := &timerMock{}
	w := &bytes.Buffer{}
	r := bytes.NewBufferString("\n")
	TimerSeconds := 30

	StartTimer(w, r, timer)

	if timer.duration != TimerSeconds {
		t.Errorf("it should set timer for %d seconds, set for %d",
			TimerSeconds, timer.duration)
	}

	if w.String() != "Ready to start?" {
		t.Errorf("it should ask user if the user is ready, got %s", w.String())
	}
}

func assertStringCalls(t *testing.T, flagger *flaggerMock) {
	t.Helper()
	if flagger.stringVarCalls != 1 {
		t.Errorf("it should call StringVar %d times, called %d",
			1, flagger.stringVarCalls)
	}
}

func assertIntCalls(t *testing.T, flagger *flaggerMock) {
	t.Helper()
	if flagger.intVarCalls != 1 {
		t.Errorf("it should call IntVar %d times, called %d",
			1, flagger.intVarCalls)
	}
}

func assertFlags(t *testing.T, flagger *flaggerMock) {
	t.Helper()

	expectedNames := []string{FileFlag, TimerFlag}
	expectedUsages := []string{FileFlagUsage, TimerFlagUsage}
	expectedStringValues := []string{FileFlagValue}
	expectedIntValues := []int{TimerFlagValue}

	if !reflect.DeepEqual(expectedNames, flagger.varNames) {
		t.Errorf("it should setup flag names to be %v, got %v",
			expectedNames, flagger.varNames)
	}

	if !reflect.DeepEqual(expectedUsages, flagger.varUsages) {
		t.Errorf("it should setup flag usages to be %v, got %v",
			expectedUsages, flagger.varUsages)
	}

	if !reflect.DeepEqual(expectedStringValues, flagger.varStringValues) {
		t.Errorf("it should setup string values to be %v, got %v",
			expectedStringValues, flagger.varStringValues)
	}

	if !reflect.DeepEqual(expectedIntValues, flagger.varIntValues) {
		t.Errorf("it should setup int values to be %v, got %v",
			expectedIntValues, flagger.varIntValues)
	}

}
