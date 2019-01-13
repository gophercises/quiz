package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("f", "problems.csv", "csv file with 2 columns, question & answer")
	timeLimit := flag.Int("t", 30, "the time limit for the quiz in seconds")
	hasShuffle := flag.Bool("s", false, "randomize question order")
	flag.Parse()

	correctCount := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Press enter to start the quiz")
	reader.ReadBytes('\n')

	b, fileErr := os.Open(*csvFilename)
	if fileErr != nil {
		fmt.Print(fileErr, "\nFailed to read the CSV file provided")
		os.Exit(1)
	}

	r := csv.NewReader(b)
	records, csvErr := r.ReadAll()

	if csvErr != nil {
		fmt.Print(csvErr, "\nFailed to parse the CSV file provided")
		os.Exit(1)
	}

	timeout := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	ansCh := make(chan string)
	getUserInput := func() {
		userAns, _ := reader.ReadString('\n')
		ansCh <- strings.TrimSpace(strings.ToLower(userAns))
	}

	if *hasShuffle {
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

qLoop:
	for i, record := range records {
		fmt.Printf("Question %d:\t%s\t", i+1, record[0])
		go getUserInput()

		select {
		case <-timeout.C:
			fmt.Println("\nToo slow!")
			break qLoop
		case ans := <-ansCh:
			if ans == record[1] {
				correctCount++
			}
		}
	}

	fmt.Printf("You scored %d out of %d!\n", correctCount, len(records))
}
