package main

import (
	"encoding/csv"
	"flag"
	_ "flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	CORRECT = iota
	INCORRECT
	TIMED_OUT
)

func scanAnswer(values chan string) {
	var inputAnswer string
	fmt.Print("Your answer: ")
	_, errScanAnswer := fmt.Scan(&inputAnswer)
	if errScanAnswer != nil {
		fmt.Println("Error while scanning answer, " + errScanAnswer.Error())
	}
	values <- inputAnswer // values <- value writing to channel
}

func timer(canceledCheck chan bool, timeOut chan bool, timeLimit int) {
	timerT := time.NewTimer(time.Duration(timeLimit) * time.Second)
	select {
	case <-timerT.C:
		timeOut <- true // writing to channel
	case canceledCheck := <-canceledCheck: // read from channel
		if canceledCheck == true {
			timerT.Stop()
		}
	}
}

func readFromCSVFile(fileName *string) ([][]string, error) {
	fileReader, errOpenFile := os.Open(*fileName)
	if errOpenFile != nil {
		fmt.Println("Error opening file: " + errOpenFile.Error())
		return [][]string{}, errOpenFile
	}
	csvReader := csv.NewReader(fileReader)
	csvReader.FieldsPerRecord = -1
	fileContent, errReadFile := csvReader.ReadAll()
	if errReadFile != nil {
		fmt.Println("Error reading from file: " + errReadFile.Error())
		return fileContent, errReadFile
	}
	return fileContent, nil
}

func handleQuestion(questionNumber int, question string, answer string, timeLimit *int) (int, error) {

	fmt.Println(fmt.Sprint(questionNumber) + ") " + fmt.Sprint(question))

	var inputAnswer string

	fmt.Print("Press Enter to continue")

	_, errScan := fmt.Scanln()
	if errScan != nil {
		fmt.Println("Error scanning new line: " + errScan.Error())
		return INCORRECT, errScan
	}

	values := make(chan string)
	canceledCheck := make(chan bool)
	timeOut := make(chan bool)

	go scanAnswer(values)
	go timer(canceledCheck, timeOut, *timeLimit)

	timedOut := false

	select {
	case timedOut = <-timeOut:
		if timedOut == true {
			fmt.Println("\nTime Out...")
		}
	case inputAnswer = <-values:
		canceledCheck <- true // write to channel to cancel the timer
	}

	if timedOut == true {
		return TIMED_OUT, nil
	}
	if inputAnswer == answer {
		return CORRECT, nil
	}
	return INCORRECT, nil
}

func handleQuizGame(questions [][]string, timeLimit *int) error {

	var correctCount int
	totalCount := len(questions)

	for count, line := range questions {
		fieldsCount := len(line)
		answer := line[fieldsCount-1]
		questionLine := line[:fieldsCount-1]
		question := strings.Join(questionLine, ",")

		status, errHandleQuestion := handleQuestion(count, question, answer, timeLimit)
		if errHandleQuestion != nil {
			fmt.Println("Error handling question, " + errHandleQuestion.Error())
		}

		if status == TIMED_OUT {
			break
		} else if status == CORRECT {
			correctCount++
		}
	}
	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Println("Correct Answers: " + fmt.Sprint(correctCount) + "\nTotal Questions: " + fmt.Sprint(totalCount))
	fmt.Println("-------------------------------------------------------------------------------------")

	return nil
}

func main() {

	//1. Read arguments
	fileName := flag.String("QuestionBankFile", "problems.csv", "Name of CSV file to read questions and answer from")
	timeLimit := flag.Int("Time[in seconds]", 15, "Time limit for each question")
	flag.Parse()

	//1.Read CSV file
	fileContent, errReadFile := readFromCSVFile(fileName)
	if errReadFile != nil {
		fmt.Println("Error reading file, " + errReadFile.Error())
	}

	errHandleQuizGame := handleQuizGame(fileContent, timeLimit)
	if errHandleQuizGame != nil {
		fmt.Println("Error handling quiz game, " + errHandleQuizGame.Error())
	}

}
