package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/zakaria-chahboun/cute"
)

// Struct of the exercise
type Exercise struct {
	Question, Answer string
}

// Struct of the flags
type Flags struct {
	Csv     string
	Limit   int
	Shuffle bool
}

// Check my errors
func checkErr(description string, err error) {
	if err != nil {
		log.Println(description, err)
	}
}

// Save each flag result in an variable type Flag, parse it and return it's pointer.
func parsing() Flags {
	options := new(Flags)
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.StringVar(&options.Csv, "csv", "problems.csv", "Insert the filepath of the csv file.")
	flag.IntVar(&options.Limit, "limit", 30, "Insert time limit of the quiz")
	flag.BoolVar(&options.Shuffle, "shuffle", false, "Wether the order of the csv file is shuffled or not")
	flag.Parse()
	return *options
}

// Read csv all at once
func read() [][]string {
	csvFile, err := os.Open(parsing().Csv)
	checkErr("Error opening file", err)

	defer csvFile.Close() // Do not close it until finished

	r := csv.NewReader(csvFile) // New reader
	r.FieldsPerRecord = -1      // Read all fields, do not expect a fixed number
	reading, err := r.ReadAll() // Data read
	checkErr("Error reading file", err)
	return reading // Give the data read as a result
}

// From the data read, create a slice that supports values type Exercise
func serialize(data [][]string) []Exercise {
	var Exercises []Exercise
	var exercise Exercise
	for _, record := range data {
		exercise.Question = record[0]
		exercise.Answer = record[1]
		Exercises = append(Exercises, exercise)
	}
	return Exercises
}

// From an slice of type Exercise, shuffle the order of the exercises
func shuffleOrder(exercises []Exercise) {
	rand.Shuffle(len(exercises), func(i, j int) {
		exercises[i], exercises[j] = exercises[j], exercises[i]
	})
}

// Creates cute list that prints quiz results
func quizResults(correct, questions int) {
	results := cute.NewList(cute.BrightBlue, "Quiz Results!")
	results.Addf(cute.BrightGreen, "Correct Answers: %d", correct)
	results.Addf(cute.BrightRed, "Incorrect Answers: %d", questions-correct)
	results.Addf(cute.BrightPurple, "Total Questions: %d", questions)
	results.Print()
}

func main() {
	var reply string
	var correct, questions int

	timer := time.NewTimer(time.Duration(parsing().Limit) * time.Second) // Timer duration
	Exercises := serialize(read())                                       // Slice with exercises

	if parsing().Shuffle {
		shuffleOrder(Exercises) // if true, shuffle order
	}

	for i := range Exercises { // iterate over the exercises slice
		cute.Println("Exercise", Exercises[i].Question, "= ")
		answerCh := make(chan string) // create a channel where it will receive the answer
		go func() {                   // while the timer run, listen to the answer
			fmt.Scanln(&reply)
			answerCh <- reply // store the scanned reply in the channel
		}()
		select {
		case <-timer.C: // when the timer runs out
			fmt.Println("") // print new line for pretty output
			cute.Println("Time is over!")
			quizResults(correct, questions) // print cute quiz result
			os.Exit(0)                      // exit with success
		case reply := <-answerCh: // store the answer channel in the reply variable
			if strings.TrimSpace(reply) == Exercises[i].Answer { // trim spaces of the reply to avoid incorrect replies
				cute.Println("Correct!")
				correct++ // count the correct replies
			} else {
				cute.Println("Incorrect!", "Your reply:", strings.TrimSpace(reply), "Answer:", Exercises[i].Answer) // if incorrect, show the correct answer
			}
			questions++ // count the questions
		}
	}
	quizResults(correct, questions)
}
