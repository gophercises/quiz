package question

import (
	"bytes"
	"io"
	"testing"
)

func Test_quiz_Answer(t *testing.T) {
	type fields struct {
		problem       string
		answer        string
		reader        io.Reader
		write         io.Writer
		totalNumber   int
		correctNumber int
	}
	type want struct {
		correct int
		total   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"success", fields{answer: "2", write: &bytes.Buffer{}, reader: bytes.NewBufferString("2")}, true},
		{"faild", fields{answer: "2", write: &bytes.Buffer{}, reader: bytes.NewBufferString("3")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			problem := &quiz{
				problem: tt.fields.problem,
				answer:  tt.fields.answer,
				reader:  tt.fields.reader,
				write:   tt.fields.write,
			}
			got := problem.Answer()
			if got != tt.want {
				t.Errorf("Answer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quizzes_Next(t *testing.T) {
	type fields struct {
		quizzes  []Quiz
		current  int
		answered map[int]bool
		writer   io.Writer
		random   bool
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"num 5", fields{quizzes: []Quiz{NewQuiz("", ""), NewQuiz("", ""), NewQuiz("", ""), NewQuiz("", ""), NewQuiz("", "")}}, 5},
		{"random", fields{random: true, quizzes: []Quiz{NewQuiz("", ""), NewQuiz("", ""), NewQuiz("", ""), NewQuiz("", ""), NewQuiz("", "")}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			problem := &quizzes{
				quizzes:  tt.fields.quizzes,
				current:  tt.fields.current,
				answered: tt.fields.answered,
				writer:   tt.fields.writer,
				random:   tt.fields.random,
			}
			total := 0
			for ; err == nil; err = problem.Next() {
				total++
			}
			if (total - 1) != tt.want {
				t.Errorf("Next() = %v, want %v", total-1, tt.want)
			}
		})
	}
}
