package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"os"
	"math/rand"
	"time"
	"strings"
)

// Create a struct to represent a question model, which holds a question and an answer
type Question struct {
	question string
	answer string	
}

// Create a struct to represent a quiz model which holds a set of questions and a score
type Quiz struct{
	score int64
	questions []Question
}

func takeQuiz(quiz *Quiz, timeout chan string){
	/*
		This function is a function that is used to take a quiz
		it takes a pointer to a quiz object as a parameter and it 
		runs the quiz operation, when the quiz is done send done message to the timeout chanel
	*/
	for _, question := range quiz.questions {
		var answer string
		fmt.Println(question.question)
		fmt.Scan(&answer)
		// if an answer is provided and the answer in the question struct match then the answer is correct
		// remove white spaces and make strings lower case
		if strings.TrimSpace(strings.ToLower(answer)) == strings.TrimSpace(strings.ToLower(question.answer)) {
			quiz.score ++
		}
	}
	timeout <- "Done!"
}

func generateScore(quiz Quiz){
	/*
		This is a function that is used to generate a score report from 
		a quiz object
	*/
	num_questions := len(quiz.questions)
	fmt.Println("There were", num_questions, "Questions")
	fmt.Println("You scored", quiz.score, "/", num_questions)

	// Calculate the average of the quiz that was taken
	average := (float64(quiz.score) / float64(num_questions)) * float64(100)
	fmt.Println("You're average is :", average, "%")
}

var (
	name string
	seconds int
	num_q int
	csv_file string
)

func init() {
	// Initialize all the flags to be used in the program and parse them
	flag.StringVar(&name, "name", "kosi", "the name of the person using this application")
	flag.IntVar(&seconds, "time", 10, "the amount of time given to the user to take the quiz")
	flag.IntVar(&num_q, "num_q", 5, "the number of questions that should be given to the user")
	flag.StringVar(&csv_file, "csv_file", "problems.csv", "specify the name of the csv file to be used")
	flag.Parse()
}

func main() {

	fmt.Println("Hello", name)
	// Read from the specified file and store it in a variable called file
	file, err := os.Open(csv_file)
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to open file because of error: ", err)
		return 
	}
	// Convert file to csv
	csv_data := csv.NewReader(file)

	if err != nil {
		fmt.Println("Failed to convert file into csv")
		return
	}
	// read rows from csv and store it in a variable called records
	records, err := csv_data.ReadAll()
	if err != nil {
		fmt.Println("Unable to read data from csv")
		return
	}

	if num_q > len(records){
		fmt.Println("The questions from the csv file are less than the questions you requested for")
		return
	}
	// Create a list of questions and get the questions from the records list
	questions := make([] Question,0)
	for _, record := range records {
		question := Question{question: record[0], answer: record[1]}
		questions = append(questions, question)
	}

	// Shuffle the questions
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i] })

	// Set the question set to the length of questions specified by the user
	question_set :=questions[:num_q]
	quiz := Quiz{score: 0, questions: question_set}
	// Create channel to know when the user is done taking a quiz or when the time for the quiz has expired
	timeout := make(chan string)

	// initate timer
	go func() {
		timewatcher := time.After(time.Duration(int(seconds)) * time.Second)
		message :=<- timewatcher
		fmt.Println(message)
		timeout <- "Timeout!"
	}()

	// Take the quiz
	go takeQuiz(&quiz, timeout)
	status :=<- timeout

	// Check if user finished the quiz or the quiz was timed out
	if status == "Done!"{
		fmt.Println("Done with the Quiz")
	}
	if status == "Timeout!"{
		fmt.Println("Oops! looks like you ran out of time!")
	}
	// Generate a report of the quiz taken
	generateScore(quiz)
}
