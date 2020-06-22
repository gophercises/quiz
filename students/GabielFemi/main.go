// Create a program that will read in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how many questions they get right and how many the get incorrect. regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.'
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// methods
	// read the csv into a struct
	// render the csv to the user

// data structures
	// a slice to contain all the csv data

// operations
	// Increment the score by 1 if correct else pass on to next instruction

// how to

var (
	file		string
	)

type problem struct {
	question string
	answer   string
}

func init() {
	flag.StringVar(&file,"csv", "read a certain csv file", "problems.csv")
	flag.Parse()
}

func main() {
	csvPath, err := filepath.Abs(file)
	handleError(err)


	file, err := os.Open(csvPath)

	formattedError := fmt.Sprintf("Failed to open %s", csvPath)
	handleErrorWithMessage(err, formattedError)

	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()

	startQuiz(parseLines(csvData))
}
// Scan the letter,
// Send it to -> customercare@ikejaelectric.com
// The subject is on the letter
//

func parseLines(linesReadFromCSV [][]string) []problem {
	parsedProblem := make([]problem, len(linesReadFromCSV))
	for index, pair := range linesReadFromCSV {
		parsedProblem[index] = problem{
			question: pair[0],
			answer: pair[1],
		}
	}
	return parsedProblem
}

func startQuiz(quizPair []problem) {
	correct := 0

	for _, quiz := range quizPair {
		fmt.Print("Question -> ", quiz.question, ": ")

		var answerFromUser string

		_, err := fmt.Scanf("%s",&answerFromUser)

		handleError(err)

		if answerFromUser == quiz.answer {
			fmt.Println("Yeah!")
			correct ++
		}

		fmt.Println(answerFromUser)
	}

	fmt.Println("You got", correct, "correctly")
}


// Bufio reader

// Helper functions

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleErrorWithMessage(err error, message string) {
	if err != nil {
		log.Fatalln(message)
	}
}