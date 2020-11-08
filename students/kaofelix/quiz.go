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

func normalize(s string) string {
	return strings.Trim(strings.ToLower(s), " ")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	timeLimit := flag.Int64("time", 30, "Time limit for quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "When true, shuffle the questions")
	flag.Parse()

	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}

	csv := csv.NewReader(file)
	rows, err := csv.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to the Quiz! Time limit is %d seconds.\nPress enter to start...", *timeLimit)
	fmt.Scanln()

	var score int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		time.Sleep(time.Duration(*timeLimit) * time.Second)
		fmt.Println("Time's up!")
		wg.Done()
	}()

	go func() {
		var order []int

		if *shuffle {
			order = rand.Perm(len(rows))
		} else {
			order = make([]int, len(rows))
			for i := 0; i < len(rows); i++ {
				order[i] = i
			}
		}

		for i, n := range order {
			question := rows[n][0]
			rightAnswer := rows[n][1]

			fmt.Printf("%d) %s\n", i+1, question)

			var answer string
			fmt.Scanln(&answer)
			if normalize(answer) == normalize(rightAnswer) {
				score++
			}
		}
		fmt.Println("Finish!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Score: %d/%d\n", score, len(rows))
}
