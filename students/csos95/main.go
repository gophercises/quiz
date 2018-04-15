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
	// set flags
	csvPath = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit   = flag.Int("limit", 30, "the time limit for the quiz in seconds'")
)

func main() {
	// parse the flags
	flag.Parse()

	// open the csv file
	file, err := os.Open(*csvPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// open the csv file
	csvReader := csv.NewReader(file)
	// parse the csv file
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
		return
	}

	// put question/answer pairs into a map where questions are keys and answers are values
	qaPair := make(map[string]string, len(csvData))
	for _, data := range csvData {
		qaPair[data[0]] = data[1]
	}

	// create a ticker for the time limit and a channel to signal the user finished the quiz
	ticker := time.NewTicker(time.Second * time.Duration(*limit))
	done := make(chan bool)

	// create a scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	var userAnswer string
	qNum, numCorrect := 0, 0
	go func() {
		// iteration order for maps in go is randomized so the questions won't be in the same order every time
		for question, answer := range qaPair {
			qNum++
			// ask a question
			fmt.Printf("Problem #%d: %s = ", qNum, question)
			// get an answer
			scanner.Scan()
			userAnswer = scanner.Text()
			// trim leading and trailing whitespace
			userAnswer = strings.TrimSpace(userAnswer)
			userAnswer = strings.ToLower(userAnswer)
			answer = strings.ToLower(answer)
			// check the answer
			if answer == userAnswer {
				numCorrect++
			}
		}
		done <- true
	}()

	// select chooses the first channel with an available value
	// if done is available first, the user finished
	// if ticker is available first, the time limit has been reached
	select {
	case <-done:
	case <-ticker.C:
		fmt.Println("time's up!")
	}

	// print the results
	fmt.Printf("You scored %d out of %d.\n", numCorrect, len(qaPair))
}
