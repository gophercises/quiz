package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

// Scorecard ...
type Scorecard struct {
	Total   int
	Correct int
}

var card Scorecard

func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func initializeQuiz() (string, int) {
	filePath := flag.String("csv-file", "problems.csv", "A scv file in the format 'question,answer'")
	timeLimit := flag.Int("time-limit", 30, "A time limit per question in seconds")

	flag.Parse()
	fmt.Println("--------------------------------")
	fmt.Printf("Quiz File: %v\n", *filePath)
	fmt.Println("--------------------------------")
	if *timeLimit > 0 {
		fmt.Println("--------------------------------")
		fmt.Printf("Time Limit: %v seconds (per question)\n", *timeLimit)
		fmt.Println("--------------------------------")
	}
	return *filePath, *timeLimit
}

func main() {

	filePath, timeLimit := initializeQuiz()

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	records := csv.NewReader(f)

	if err != nil {
		log.Fatal(err)
	}

	card.Correct = 0
	card.Total = 0

	for {
		record, err := records.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		card.Total++
		duration := time.Duration(
			int64(timeLimit) * int64(time.Second),
		)
		timer := time.NewTimer(duration)

		if timeLimit > 0 {
			go func() {
				<-timer.C
				fmt.Printf("Times up. You got %v out of %v correct", card.Correct, card.Total)
				os.Exit(0)
			}()
		}

		fmt.Printf("What is %v? ", record[0])
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		stop := timer.Stop()
		if stop {

			fmt.Printf("Your Answer:    %v", answer)
			fmt.Printf("Correct Answer: %v \n", record[1])
		}

		a, err := strconv.Atoi(trimLastChar(answer))
		if err != nil {
			fmt.Println("Numbers only please.")
		}
		correct, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		if a == correct {
			fmt.Println("Correct!")
			card.Correct++
		} else {
			fmt.Println("Incorrect.")
		}
		fmt.Println("-----")

	}
	fmt.Printf("You got %v out of %v correct", card.Correct, card.Total)
}
