package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type parsedCsv = [][]string

type problem = []string

func main() {
	problemFilePath, timeout := parseCommandLineArguments()
	problemSet := parseProblemsCsv(problemFilePath)

	var correctAnswers int
	launchTimer(timeout, problemSet, &correctAnswers)
	poseProblems(problemSet, &correctAnswers)
	printResults(problemSet, &correctAnswers)
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		return true
	}
}

func verifyCsvFitsProblemsFormat(csvContent parsedCsv) {
	for _, row := range csvContent {
		if len(row) != 2 {
			panic("The CSV seemed to be malformatted. We require to receive a CSV with 2 columns for every row.")
		}
	}
}

func parseProblemsCsv(csvLocation string) parsedCsv {
	if !fileExists(csvLocation) {
		panic(fmt.Sprintf("The file \"%s\" does not exist. You need a CSV file either called problems.csv or you hand us a location pointing to a different CSV file matching the problem format.", csvLocation))
	}
	csvContent, err := ioutil.ReadFile(csvLocation)
	if err != nil {
		panic(fmt.Sprintf("Although the problems.csv file \"%s\" existed, we had trouble reading it.", csvLocation))
	}
	stringifiedCsv := string(csvContent)
	csvReader := csv.NewReader(strings.NewReader(stringifiedCsv))
	problemSet, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	verifyCsvFitsProblemsFormat(problemSet)
	return problemSet
}

func poseProblem(problemItem problem) bool {
	question := problemItem[0]
	answer := problemItem[1]
	fmt.Println(fmt.Sprintf("Q: %s", question))

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	// convert CRLF to LF
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	if userInput == answer {
		fmt.Println("Correct")
		return true
	} else {
		fmt.Println("Incorrect")
		return false
	}
}

func poseProblems(problemSet parsedCsv, correctAnswers *int) {
	for _, problem := range problemSet {
		correctAnswerGiven := poseProblem(problem)
		if correctAnswerGiven {
			*correctAnswers++
		}
	}
}

func printResults(problemSet parsedCsv, correctAnswers *int) {
	totalQuestions := len(problemSet)
	fmt.Println()
	fmt.Println(fmt.Sprintf("Total questions: %d", totalQuestions))
	fmt.Println(fmt.Sprintf("Correct answers: %d", *correctAnswers))
	fmt.Println(fmt.Sprintf("Incorrect answers: %d", totalQuestions-*correctAnswers))
}

func launchTimer(timeout int, problemSet parsedCsv, correctAnswers *int) {
	timer1 := time.NewTimer(time.Duration(timeout) * time.Second)
	go func() {
		<-timer1.C
		fmt.Println()
		fmt.Println("Time ran out!")
		printResults(problemSet, correctAnswers)
		os.Exit(0)
	}()
}

func parseCommandLineArguments() (string, int) {
	problemsPtr := flag.String("problems", "problems.csv", "A relative file path pointing to the CSV set of problems you want to ask the user.")
	timeoutPtr := flag.Int("timeout", 10, "An integer. The number of seconds the user has time to answer all the questions.")
	flag.Parse()
	return *problemsPtr, *timeoutPtr
}
