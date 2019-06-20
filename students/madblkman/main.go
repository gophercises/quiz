package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Grab the problems file from the cli
	problemsFile := os.Args[1]
	var rightAnswers int
	var wrongAnswers int
	var total int

	// Open the csv file
	inputBytes, err := ioutil.ReadFile(problemsFile)
	if err != nil {
		fmt.Print(err)
	}

	// Convert the problems file from bytes to a string
	data := string(inputBytes)

	// Create a Reader type from the converted data
	r := csv.NewReader(strings.NewReader(data))
	r.Comma = ','

	// Read all of the records into a variable
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for _, record := range records {
		fmt.Printf("What is %s?\n", record[0])
		text, _ := reader.ReadString('\n')
		if text == record[1] {
			rightAnswers++
			fmt.Printf("Input: %sAnswer: %s\n", text, record[1])
		} else {
			wrongAnswers++
			fmt.Printf("Input: %sAnswer: %s\n", text, record[1])
		}
		total++
	}

	fmt.Printf("Right: %d\nWrong: %d\nTotal: %d\n", rightAnswers, wrongAnswers, total)
}
