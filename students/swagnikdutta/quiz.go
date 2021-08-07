package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filepath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30)")
	flag.Parse()

	file, err := os.Open(*filepath)

	if err != nil {
		log.Fatal("Error while reading csv\n", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	file.Close()

	if err != nil {
		log.Fatal("There was an error reading records from file\n", err)
	}

	correct := 0
	input := bufio.NewScanner(os.Stdin)
	for idx, record := range records {
		// Here two things can happen
		// Either the timer will expire before the user can answer or the user will answer before the timer expires.
		// In either case, I need to listen for both sorts of events.
		// So there has to be 2 channels to listen to:
		//   	First, through which the user's answer will be received.
		//		Second, through which the timer will notify that the time is up
		//
		// There are two concurrent processes involved.
		// 1. The problem statement has to be shown to the user. We wait for or handle the user's response.
		//    This part has to be written inside a goroutine to make it run concurrently.
		// 2. Immediately after showing the problem statement, we start the timer.

		answerChannel := make(chan string)

		// Step 1: Showing question and prompting user for answer
		go func() {
			fmt.Printf("Problem #%d: %s = ", idx+1, record[0])
			input.Scan()
			userAnswer := input.Text()
			// Answer provided by the user has to be sent to the main goroutine (from this goroutine) via a channel.
			// So we need to have a channel created in the first place - answerChannel (created above)
			answerChannel <- userAnswer
		}()

		// Step 2: starting the timer
		timer := time.NewTimer(time.Duration(*limit) * time.Second)

		// Wait and listen for either events
		select {
		case <-timer.C:
			// When the timer expires, the current time is sent to channel C
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(records))
			return
		case userAnswer := <-answerChannel:
			if userAnswer == record[1] {
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d.\n", correct, len(records))
}
