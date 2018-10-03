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

func quiz(records []Quiz) (*Stat, error) {
	var stat Stat
	reader := bufio.NewReader(os.Stdin)
	for _, elem := range records {
		fmt.Print(elem.question, ":")
		ans, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		stat.all++
		if strings.TrimRight(ans, "\r\n") == elem.answer {
			stat.correct++
		} else {
			stat.incorrect++
		}
	}
	return &stat, nil
}

func main() {
	var f string
	flag.StringVar(&f, "f", "problems.csv", "input file in csv format")
	flag.Parse()
	recs, err := readCSV(f)
	if err != nil {
		log.Fatal(err)
	}
	stat, err := quiz(recs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Question answered: %v, Correct: %v, Incorrect: %v\n", stat.all, stat.correct, stat.incorrect)
}
