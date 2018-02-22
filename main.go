package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type item struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	csvFile, err := os.Open("problems.csv")

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

	resultsJSON, err := json.MarshalIndent(quiz, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resultsJSON))
}
