package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readCsv(filePath string, shuffle bool) [][]string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, ".\n", err)
	}

	r := csv.NewReader(file)

	content, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to read supplied .csv file at "+filePath, ".\n", err)
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(content), func(i, j int) { content[i], content[j] = content[j], content[i] })
	}

	return content
}

func readAnswer() string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func runQuiz(qs [][]string, timeLimit int) (int, int) {

	score := 0
	numQs := len(qs)

	// create channels
	answerChannel := make(chan string)
	timerChannel := make(chan bool, 1)

	for q := 0; q < numQs; q++ {

		question := qs[q][0]
		correctAnswer := qs[q][1]

		// Start timer
		timer := time.NewTimer(time.Second * time.Duration(timeLimit))

		go func() { // execute timer process
			<-timer.C
			timerChannel <- true
		}()

		go func() { // execute answer read process
			fmt.Printf("Problem #%s: %s = ", strconv.Itoa(q+1), question)
			answerChannel <- readAnswer()
		}()

		select { // case with first channel to receive values gets executed
		case answer := <-answerChannel:
			if answer == correctAnswer {
				score += 1
			}
			timer.Stop()
		case <-timerChannel:
			return score, numQs
		}
	}
	return score, numQs
}

func main() {

	// get command line args
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shufflePtr := flag.Bool("shuffle", false, "shuffle the quiz order")
	flag.Parse()

	// read CSV
	csv := readCsv(*csvPtr, *shufflePtr)

	// run Quiz
	score, numQs := runQuiz(csv, *limitPtr)
	fmt.Printf("\nYou scored %d out of %d.", score, numQs)
}
