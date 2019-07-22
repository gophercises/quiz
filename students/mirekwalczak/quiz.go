package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// Quiz is structure for questions and answers
type Quiz struct {
	question, answer string
}

// Stat is struct for quiz statistics
type Stat struct {
	all, correct, incorrect int
}

func readCSV(file string) ([]Quiz, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var quizes []Quiz
	r := csv.NewReader(f)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		quizes = append(quizes, Quiz{
			strings.TrimSpace(line[0]),
			strings.TrimSpace(line[1]),
		})
	}
	return quizes, nil
}

func quiz(records []Quiz, timeout int) (*Stat, error) {
	var stat Stat
	reader := bufio.NewReader(os.Stdin)

	timer := time.NewTimer(time.Second * time.Duration(timeout))
	errs := make(chan error)

	go func() {
		for _, quiz := range records {
			fmt.Print(quiz.question, ":")

			ans, err := reader.ReadString('\n')
			if err != nil {
				errs <- err
			}
			stat.all++
			if strings.TrimRight(ans, "\r\n") == quiz.answer {
				stat.correct++
			} else {
				stat.incorrect++
			}
		}
	}()

	select {
	case <-errs:
		return nil, <-errs
	case <-timer.C:
		fmt.Println("\ntime's up!")
	}

	return &stat, nil
}

func main() {
	f := flag.String("f", "problems.csv", "input file in csv format")
	t := flag.Int("t", 30, "timeout for the quiz, in seconds")
	flag.Parse()
	recs, err := readCSV(*f)
	if err != nil {
		log.Fatal(err)
	}
	stat, err := quiz(recs, *t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nQuestion answered: %v, Correct: %v, Incorrect: %v\n", stat.all, stat.correct, stat.incorrect)
}
