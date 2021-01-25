package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	nDone  int
	nRight int
)

/*
 * If expres is true, call log.Fatalln() to exit the program
 */
func check(expres bool, err error) {
	if expres {
		log.Fatalln(err)
	}
}

/*
 * Get the filename from the command-line argument, and the time the program was run
 */
func readArguments() (string, time.Duration) {
	fileName := flag.String("f", "problems.cvs", "CSV File that conatins quiz questions")
	timeLimit := flag.Duration("t", 30*time.Second, "Time Limit for each question")
	flag.Parse()

	return *fileName, *timeLimit
}

/*
 * Read the file in CSV format
 */
func readCSVFile(fileName string) ([][]string, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

/*
 * Input
 */
func getInput(input chan<- string) {
	in := bufio.NewReader(os.Stdin)

	for {
		result, err := in.ReadString('\n')
		check(err != nil, err)

		input <- strings.Trim(strings.ToLower(result), " ")
	}
}

func init() {
	log.SetFlags(log.LstdFlags ^ log.Ldate)
}

func main() {
	fileName, timeLimit := readArguments()

	fmt.Print("Are you ready?")
	fmt.Scanln()

	// Timed out
	timeOut := time.NewTicker(timeLimit)
	/*
	 * timeOut := make(chan bool)
	 * time.AfterFunc(timeLimit, func() {
	 *     timeOut <- true
	 * })
	 */
	records, err := readCSVFile(fileName)
	check(err != nil, err)

	userAnswerChan := make(chan string)
	go getInput(userAnswerChan)

LOOP:
	for i := range records {
		question, answer := records[i][0], records[i][1]
		fmt.Printf("Que is %s ? ", question)

		select {
		case <-timeOut.C:
			break LOOP
		case usernswer, ok := <-userAnswerChan:
			nDone++
			if ok && usernswer == answer {
				nRight++
			}
		}
	}

	fmt.Printf("\nYou answered %d out of %d questions correctly.\n", nDone, nRight)
}
