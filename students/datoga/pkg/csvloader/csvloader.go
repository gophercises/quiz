package csvloader

import (
	"encoding/csv"
	"io"

	"github.com/datoga/quiz/pkg/model"
)

type CSVLoader struct {
	in io.Reader
}

func NewCSVLoader(input io.Reader) *CSVLoader {
	return &CSVLoader{in: input}
}

func (loader CSVLoader) CSVToModel() (model.Quiz, error) {
	quiz := []model.QuizItem{}

	r := csv.NewReader(loader.in)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return quiz, err
		}

		quizItem := model.QuizItem{
			Question: record[0],
			Solution: record[1],
		}

		quiz = append(quiz, quizItem)
	}

	return quiz, nil
}
