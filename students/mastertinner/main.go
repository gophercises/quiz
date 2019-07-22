package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// problem is a problem to be solved by a user.
type problem struct {
	challenge     string
	correctAnswer string
}

// scoreboard keeps track of a user's score.
type scoreboard struct {
	total   int64
	correct int64
}

func main() {
	var (
		csvFileName = flag.String("csv", "problems.csv", "the location of the CSV file")
		timeLimit   = flag.Int("time-limit", 30, "the time limit for the user to answer all problems")
		doShuffle   = flag.Bool("shuffle", false, "whether to shuffle the order of problems or not")
	)
	flag.Parse()

	problems, err := readProblemsFromCSVFile(*csvFileName)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading problems from CSV file: %s", err))
	}
	if *doShuffle {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

	doneCh := make(chan bool)
	scr := &scoreboard{
		total: int64(len(problems)),
	}

	go func() {
		for i, p := range problems {
			fmt.Printf("What is %s?\n", p.challenge)
			inputReader := bufio.NewReader(os.Stdin)
			fmt.Print("Answer: ")
			answer, err := inputReader.ReadString('\n')
			if err != nil {
				log.Fatal(fmt.Errorf("error reading user input: %s", err))
			}
			if purifyString(answer) == purifyString(p.correctAnswer) {
				scr.correct++
				fmt.Println("You are correct!")
			} else {
				fmt.Printf("Unfortunately not... The correct answer is %s\n", p.correctAnswer)
			}
			fmt.Printf("Your current score is %v/%v\n\n", scr.correct, scr.total)
			if i == len(problems)-1 {
				doneCh <- true
			}
		}
	}()

	go func() {
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
		<-timer.C
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Your time is up...")
		doneCh <- true
	}()

	<-doneCh
	fmt.Println("")
	fmt.Printf("Your final score is %v/%v\n\n", scr.correct, scr.total)
}

// readProblemsFromCSVFile reads problems from a CSV file.
func readProblemsFromCSVFile(fileName string) ([]problem, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []problem
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("error reading CSV line: %s", err)
		}
		if len(line) != 2 {
			return nil, errors.New("invalid line in CSV")
		}
		p := problem{
			challenge:     line[0],
			correctAnswer: line[1],
		}
		problems = append(problems, p)
	}
	return problems, nil
}

// purifyString strips all unneeded variations from a string.
func purifyString(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}
