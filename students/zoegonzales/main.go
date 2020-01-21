package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// flags for csv file, time limit, and shuffling questions
	fn := flag.String("file", "problems.csv", "a csv file")
	seconds := flag.Int64("seconds", 30, "quiz timer")
	shuffled := flag.Bool("shuffle", false, "shuffle quiz questions")
	flag.Parse()
	fmt.Printf("Press enter to play. Once the quiz is started you will have %v seconds to answer all of the questions.", *seconds)
	fmt.Scanln()
	if *shuffled {
		startQuiz(shuffle(parseCSV(*fn)), *seconds)
	} else {
		startQuiz(parseCSV(*fn), *seconds)
	}
}

func parseCSV(file string) [][]string {
	// open and read file, return data as slice of slices
	f, openError := os.Open(file)
	if openError != nil {
		fmt.Printf("%v", openError)
	}
	r := csv.NewReader(f)
	records, readError := r.ReadAll()
	if readError != nil {
		fmt.Printf("%v", readError)
	}
	return records
}

func startQuiz(questions [][]string, limit int64) {
	var correct int = 0
	var incorrect int = 0
	var count int = 0
	// set timer for quiz limit
	t := time.NewTimer(time.Duration(limit) * time.Second)
	defer t.Stop()
	// establish wait group; wait for all question goroutines to finish before ending main goroutine
	var wg sync.WaitGroup

	for count < len(questions) {
		// create channel for current question and add to wait group
		answered := make(chan bool)
		wg.Add(1)
		defer wg.Done()

		go func() {
			// prompt questions and scan input
			fmt.Println(questions[count][0])
			var in string
			fmt.Scanln(&in)
			// check for correct answer
			modified := strings.ToLower(strings.TrimSpace(in))
			if modified == questions[count][1] {
				correct++
			} else {
				incorrect++
			}
			count++
			answered <- true
		}()

		select {
		// current answered channel closed, continue quiz
		case <-answered:
			continue
		// timer channel closed, end quiz
		case <-t.C:
			unanswered := len(questions) - count
			fmt.Println("\nTime's up! Your results are below.")
			fmt.Printf("Correct answers: %v | Incorrect answers: %v | Unanswered questions: %v\n", correct, incorrect, unanswered)
			os.Exit(1)
		}
	}
	// user completed quiz before time was up
	fmt.Println("Congrats, you finished the quiz! Your results are below.")
	fmt.Printf("Correct answers: %v | Incorrect answers: %v\n", correct, incorrect)
	os.Exit(1)
	wg.Wait()
}

func shuffle(questions [][]string) [][]string {
	// generate new source for random numbers
	r := rand.New(rand.NewSource(time.Now().Unix()))
	// iterate over permutation of the random numbers in questions
	for i, q := range r.Perm(len(questions)) {
		// swap current item's original index with new pseudo-random index
		questions[i], questions[q] = questions[q], questions[i]
	}
	return questions
}
