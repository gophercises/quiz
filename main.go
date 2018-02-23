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

type item struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	filepath := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	csvFile, err := os.Open(*filepath)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	var quiz []item

	for {
		var line []string
		line, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		if len(line) != 2 {
			log.Fatalf("expected line to have 2 elements, got %d for %v", len(line), line)
		}

		quiz = append(quiz, item{
			Question: line[0],
			Answer:   line[1],
		})
	}

	possibleScore := len(quiz)
	actualScore := 0

	for i, item := range quiz {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Problem #%d: %s = ", i, item.Question)
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(text) == item.Answer {
			actualScore++
		}
	}

	fmt.Printf("You scored %d out of %d.", actualScore, possibleScore)
}
