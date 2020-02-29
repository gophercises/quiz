package question

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

func SliceQuestion(slice []string) (Quiz, error) {
	if len(slice) != 2 {
		return nil, errors.New("The question and answer format is not correct.")
	}
	return NewQuiz(slice[0], slice[1]), nil
}

func CSVQuizzes(path string) (Quizzes, error) {
	quizzes := make([]Quiz, 0)
	file, err := os.Open(path)
	if err != nil {
		return NewQuizzes(quizzes...), err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	for {
		line, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return NewQuizzes(quizzes...), err
		}
		q, err := SliceQuestion(line)
		if err != nil {
			return NewQuizzes(quizzes...), err
		}
		quizzes = append(quizzes, q)
	}
	return NewQuizzes(quizzes...), nil
}
