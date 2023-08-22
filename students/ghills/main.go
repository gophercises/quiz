package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var fileName string
	var timeOutSeconds int
	var shuffle bool

	// flags declaration using flag package
	flag.StringVar(&fileName, "filename", "problems.csv", "Specify filename. Default is problems.csv")
	flag.IntVar(&timeOutSeconds, "time-limit", 30, "Specify time limit in seconds. Default is 30")
	flag.BoolVar(&shuffle, "shuffle", false, "Indicates the problems should be shuffled. Default is false")

	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	questions, _ := reader.ReadAll()

	questionCount := len(questions)

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(questionCount, func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	// signal channel in case timeout is reached prior to quiz completion
	timeoutChan := make(chan bool)

	// run timer concurrent to the quiz
	go runTimer(timeOutSeconds, timeoutChan)
	correct := runQuiz(questions, timeoutChan)

	fmt.Printf("%v/%v -- %.0f%%", correct, questionCount, (float64(correct)/float64(questionCount))*100)
}

// run a timer for the indicated number of seconds and fire a value on the timeout channel
// when it ends
func runTimer(timeoutSeconds int, timeoutChan chan bool) {
	timer := time.NewTimer(time.Duration(timeoutSeconds) * time.Second)
	<-timer.C
	timeoutChan <- true
}

func runQuiz(problems [][]string, timeoutChan chan bool) (correct int) {
	numCorrect := 0
	inputChan := make(chan string)

problems:
	for _, p := range problems {
		go promptForInput(p[0], inputChan)
		// allow the timeout to fire while waiting for user input here
		select {
		case response := <-inputChan:
			if strings.TrimSpace(strings.ToLower(response)) == strings.TrimSpace(strings.ToLower(p[1])) {
				numCorrect += 1
			}
		case <-timeoutChan:
			// on timeout it is sitting on a prompt and needs
			// to be moved to next line for results display
			fmt.Println("")
			// break out of the problems loop so the control returns as the quiz is ending (incomplete)
			break problems
		}
	}

	return numCorrect
}

func promptForInput(prompt string, responseChan chan string) {
	fmt.Print(prompt, ": ")
	var response string
	n, err := fmt.Scanln(&response)
	if err == nil && n > 0 {
		responseChan <- response
	} else {
		responseChan <- ""
	}
}
