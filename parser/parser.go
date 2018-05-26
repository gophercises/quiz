package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fedepaol/quiz/quiz"
)

func ParseFile(filename string) (res quiz.Quiz, e error) {
	file, err := os.Open(filename)
	if err != nil {
		e = err
		return
	}
	return parseQuestions(file)
}

func parseQuestion(row []string) (quiz.Question, error) {
	if len(row) < 2 {
		return quiz.Question{}, fmt.Errorf("not enough values for a question %v", row)
	}

	question := strings.Join(row[:len(row)-1], ",")
	answer := row[len(row)-1]

	return quiz.Question{question, answer}, nil
}

// ParseQuestions returns a quiz based on a csv reader.
func parseQuestions(reader io.Reader) (res quiz.Quiz, e error) {
	r := csv.NewReader(reader)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			e = err
			return
		}

		q, err := parseQuestion(record)
		if err != nil {
			e = err
			return
		}

		res.AddQuestion(q)
	}

	return
}
