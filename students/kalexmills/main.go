package main

import (
"encoding/csv"
"flag"
"fmt"
"log"
"os"
"time"
)

var (
	filename  string
	timelimit time.Duration
)

func init() {
	flag.StringVar(&filename, "in", "problems.csv", "filename where problems are read from")
	flag.DurationVar(&timelimit, "time", 30*time.Second, "time to allot for quiz")
}


// problem represents a prompt and the correct response
type problem struct {
	prompt   string
	response string
}

// printResults prints a prompt to the user explaining the results achieved.
func printResults(nWrong, nTotal int) {
	nRight := nTotal - nWrong
	fmt.Printf("You answered %d out of %d questions correctly.\n", nRight, nTotal)
}

// readFile reads from the file, sending back a slice of problems.
func readFile() []problem {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file: %s", err.Error())
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// Construct a slice of problems for later consumption
	result := make([]problem, len(records))
	var i int
	for idx, rec := range records {
		if len(rec) < 2 {
			fmt.Println("Ignoring faulty or missing record from line %d", idx)
			continue
		}
		result[i] = problem{rec[0], rec[1]}
		i++
	}

	return result
}

func main() {
	flag.Parse()

	problems := readFile()

	var nWrong, nDone int
	// timeout is used to signal that the timer has elapsed
	timeout := make(chan bool, 2)
	timer := time.AfterFunc(timelimit, func() {
		timeout <- true // signal that a timeout occurred
	})

	startTime := time.Now()
	// Ask each question and get a response
	for _, prob := range problems {
		fmt.Printf("%s\n > ", prob.prompt)

		// A separate goroutine must be used to accept input from each question. At first I tried to just read from
		// Stdin on the main thread. That was a mistake. If you want to be able to interrupt the prompt when the time
		// expires, you need to have your main thread waiting on a response from either the timer or the answer thread.
		//
		// Lesson is to treat user input as an asynchronous message!
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timeout:
			fmt.Println("\nTimes up!")
			nWrong += len(problems) - nDone // All remaining problems are marked wrong
			printResults(nWrong, len(problems))
			os.Exit(0)
		case response := <-answerCh:
			nDone++
			if prob.response != response {
				nWrong++
			}
		}
	}

	if timer.Stop() {
		fmt.Printf("You completed the quiz with %s remaining\n", (timelimit - time.Since(startTime)).String())

		printResults(nWrong, len(problems))
	}
}
