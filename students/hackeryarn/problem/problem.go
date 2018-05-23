package problem

import (
	"fmt"
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
func (p Problem) CheckAnswer(r io.Reader) bool {
	answer := readAnswer(r)

	if answer != p.answer {
		return false
	}
	return true
}

func readAnswer(r io.Reader) (answer string) {
	_, err := fmt.Fscanln(r, &answer)
	if err != nil {
		log.Fatalln("Error reading in answer", err)
	}

	return strings.TrimSpace(answer)
}

// AskQuestion prints out the question
func (p Problem) AskQuestion(w io.Writer) {
	_, err := fmt.Fprintf(w, "%s: ", p.question)
	if err != nil {
		log.Fatalln("Could not ask the question", err)
	}
}

// New creates a Problem from a provided CSV record
func New(record []string) Problem {
	return Problem{
		question: record[0],
		answer:   record[1],
	}
}
