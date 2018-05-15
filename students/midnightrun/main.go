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
	"math/rand"
)

type Problem struct {
	Questions string
	Answer    string
}

func main() {
	filePath := flag.String("csv", "problems.csv", "Path to the problems file")
	limit := flag.Int("limit", 5, "Limits time to answer a question to x seconds")
	shuffle := flag.Bool("shuffle", false, "Shuffles the questions within the problems file")
	flag.Parse()

	csvFile, error := os.Open(*filePath)
	if error != nil {
		log.Fatal(error)
	}
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []Problem
	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		problems = append(problems, Problem{
			Questions: line[0],
			Answer:    strings.Trim(line[1], " "),
		})
	}

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i,j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	reader := bufio.NewReader(os.Stdin)

	correctAnswers := 0

L:
	for _, quiz := range problems {
		fmt.Printf("Question: " + quiz.Questions + " ")

		channel := make(chan string, 1)
		go func() {
			text, _ := reader.ReadString('\n')
			channel <- text
		}()

		select {
		case answer := <-channel:
			if strings.ToLower(strings.TrimRight(answer, "\n")) == strings.ToLower(quiz.Answer) {
				correctAnswers = correctAnswers + 1
			}
		case <-time.After(time.Duration(*limit) * time.Second):
			println("Times up!")
			break L
		}
	}

	fmt.Printf("You answer %d correct of %d questions", correctAnswers, len(problems))
}