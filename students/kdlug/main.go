package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "csv file with questions and answers")
	duration := flag.Int("time", 30, "quiz time limit")
	randomize := flag.Bool("random", false, "randomize questions")

	// Parse
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	questions := loadRecordsFromCsv(*csvFile)
	correct := 0
	total := len(questions) - 1

	if *randomize {
		questions = shuffle(questions)
	}

	fmt.Println("Total Questions:", total)
	fmt.Println("Duration [s]:", *duration)

	done := make(chan bool, 1)

	go func() {
		for i := 0; i < total; i++ {
			fmt.Printf("Question #%d %s = ", i+1, questions[i][0])

			answer, _ := reader.ReadString('\n')
			// convert CRLF to LF
			answer = strings.Replace(answer, "\n", "", -1)
			answer = strings.ToLower(answer)
			answer = strings.TrimSpace(answer)

			// compare answer
			if strings.Compare(questions[i][1], answer) == 0 {
				correct++
			}

		}
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Good Job!")

	case <-time.After(time.Duration(*duration) * time.Second):
		fmt.Println("\nYou reached maximum time.")
	}

	fmt.Println("Your score:", correct, "/", total)
}

// shuffle questions
func shuffle(questions [][]string) [][]string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range questions {
		np := r.Intn(len(questions) - 1)
		questions[i], questions[np] = questions[np], questions[i]
	}

	return questions
}

func loadRecordsFromCsv(csvFile string) [][]string {

	// load csv file into memory, returns bytes
	content, err := ioutil.ReadFile(csvFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(content)) // if we have string instead of btes we can use strings.NewReader(content)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records[1:len(records)]
}
