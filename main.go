package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	question	string
	answer		string
}
func main() {
	// define flags (this is a string type flag)
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer'")

	// parse all defined flags
	flag.Parse()

	// compile all defined flags
	_ = csvFilename
	file, err := os.Open(*csvFilename)
	if err!= nil {
		log.Fatalln("Failed to open the csv file:", *csvFilename)
	}

	// return a new reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Println("Failed to parse csv file")
	}
	problems := parseLInes(lines)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var answer string

		// scan the answer from the user
		// scan removes white spaces from what it receives so it won't be a good use case for string answers
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

// parse the lines from the lines array into a slice of problems
func parseLInes(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return  problems
}
// you can use the -h flag or --help flag to show all the available flags a program has
// the definition you put in usage, would be displayed to the user