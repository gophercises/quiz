package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	problemsFile string
	timeout      int
	correct      int
	b            bool
	inChan       chan int
	outChan      chan int
)

func getAnswer(solution string) {
	var inp string
	fmt.Scanln(&inp)
	if sanitize(inp) == sanitize(solution) {
		outChan <- <-inChan + 1
	} else {
		outChan <- <-inChan
	}
}

func updateCorrect() bool {
	select {
	case res := <-outChan:
		correct = res
		return true
	case <-time.After(time.Duration(timeout) * time.Second):
		close(outChan)
		fmt.Println()
		return false
	}
}

func sanitize(s string) string {
	return strings.ToLower(strings.Trim(s, "\n\r\t "))
}

func main() {
	flag.StringVar(&problemsFile, "path", "problems.csv", "This is the flag to the CSV containing the problems for the quiz")
	flag.IntVar(&timeout, "timeout", 30, "The amount of time you have for a single question in seconds")
	flag.Parse()

	file, err := os.Open(problemsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	problems, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	inChan, outChan = make(chan int, 1), make(chan int, 1)
	for c, q := range problems {
		fmt.Printf("Question %v: %v -> ", c, q[0])

		go getAnswer(q[1])
		inChan <- correct

		if !updateCorrect() {
			break
		}

	}
	fmt.Println("You got", correct, "out of", len(problems), "correct.")
}
