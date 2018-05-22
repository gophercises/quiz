package problem

import (
	"bufio"
	"log"
	"strings"
)

// Problem represents a single question answer pair
type Problem struct {
	question string
	answer   string
}

// CheckAnswer checks the answer against the provided input
func (p Problem) CheckAnswer(scanner *bufio.Scanner) bool {
	answer := readAnswer(scanner)

	if answer != p.answer {
		return false
	}
	return true
}

func readAnswer(scanner *bufio.Scanner) string {
	ok := scanner.Scan()
	if !ok {
		log.Fatalln("Error reading in answer")
	}

	return strings.TrimSpace(scanner.Text())
}

// New creates a Problem from a provided CSV record
func New(record []string) Problem {
	return Problem{
		question: record[0],
		answer:   record[1],
	}
}
