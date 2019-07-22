package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	quiz "github.com/gophercises/quiz/students/hackeryarn/myquiz"
	"github.com/gophercises/quiz/students/hackeryarn/problem"
)

const (
	// FileFlag is used to set a file for the questions
	FileFlag = "file"
	// FileFlagValue is the value used when no FileFlag is provided
	FileFlagValue = "problems.csv"
	// FileFlagUsage is the help string for the FileFlag
	FileFlagUsage = "Questions file"

	// TimerFlag is used for setting a timer for the quiz
	TimerFlag = "timer"
	// TimerFlagValue is the value used when no TimerFlag is provided
	TimerFlagValue = 30
	// TimerFlagUsage is the help string for the TimerFlag
	TimerFlagUsage = "Amount of seconds the quiz will allow"
)

// Flagger configures the flags used
type Flagger interface {
	StringVar(p *string, name, value, usage string)
	IntVar(p *int, name string, value int, usage string)
}

type quizFlagger struct{}

func (q *quizFlagger) StringVar(p *string, name, value, usage string) {
	flag.StringVar(p, name, value, usage)
}

func (q *quizFlagger) IntVar(p *int, name string, value int, usage string) {
	flag.IntVar(p, name, value, usage)
}

// Timer is used to start a timer
type Timer interface {
	NewTimer(d time.Duration) *time.Timer
}

type quizTimer struct{}

func (q quizTimer) NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}

// ReadCSV parses the CSV file into a Problem struct
func ReadCSV(reader io.Reader) quiz.Quiz {
	csvReader := csv.NewReader(reader)

	problems := []problem.Problem{}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error reading CSV:", err)
		}

		problems = append(problems, problem.New(record))
	}

	return quiz.New(problems)
}

// TimerSeconds is the amount of time allowed for the quiz
var TimerSeconds int
var file string

// ConfigFlags sets all the flags used by the application
func ConfigFlags(f Flagger) {
	f.StringVar(&file, FileFlag, FileFlagValue, FileFlagUsage)
	f.IntVar(&TimerSeconds, TimerFlag, TimerFlagValue, TimerFlagUsage)
}

// StartTimer begins a timer once the user provides input
func StartTimer(w io.Writer, r io.Reader, timer Timer) *time.Timer {
	fmt.Fprint(w, "Ready to start?")
	fmt.Fscanln(r)

	return timer.NewTimer(time.Second * time.Duration(TimerSeconds))
}

func init() {
	flagger := &quizFlagger{}
	ConfigFlags(flagger)

	flag.Parse()
}

func main() {
	file, err := os.Open(file)
	if err != nil {
		log.Fatalln("Could not open file", err)
	}

	quiz := ReadCSV(file)

	timer := StartTimer(os.Stdout, os.Stdin, quizTimer{})
	go func() {
		<-timer.C
		fmt.Println("")
		quiz.PrintResults(os.Stdout)
		os.Exit(0)
	}()

	quiz.Run(os.Stdout, os.Stdin)
}
