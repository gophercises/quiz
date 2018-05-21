package problem

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// Problem represents a single question answer pair
type Problem struct {
	question string
	answer   string
}

// CheckAnswer checks the answer against the provided input
func (p Problem) CheckAnswer(input io.Reader) bool {
	answer := readAnswer(input)

	if answer != p.answer {
		return false
	}
	return true
}

func readAnswer(input io.Reader) string {
	reader := bufio.NewReader(input)
	answerString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error reading in answer:", err)
	}

	return strings.TrimSpace(answerString)
}

// New creates a Problem from a provided CSV record
func New(record []string) Problem {
	return Problem{
		question: record[0],
		answer:   record[1],
	}
}
