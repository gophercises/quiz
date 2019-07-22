package quiz

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/gophercises/quiz/students/hackeryarn/problem"
)

func TestNew(t *testing.T) {
	problems := sampleProblems()

	want := Quiz{problems: problems, rightAnswers: 0}
	got := New(problems)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expeted to create quiz %v got %v", want, got)
	}
}

func TestRun(t *testing.T) {
	t.Run("it runs the quiz", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		quiz := createQuiz()
		runQuiz(buffer, &quiz)

		expectedResults := 2
		results := quiz.rightAnswers

		if expectedResults != results {
			t.Errorf("expected right answers of %v, got %v",
				expectedResults, results)
		}

		expectedOutput := "7+3: 1+1: You got 2 questions right!\n"

		if buffer.String() != expectedOutput {
			t.Errorf("expected full output %v, got %v",
				expectedOutput, buffer)
		}

	})
}

func sampleProblems() []problem.Problem {
	record1 := []string{"7+3", "10"}
	record2 := []string{"1+1", "2"}

	return []problem.Problem{
		problem.New(record1),
		problem.New(record2),
	}
}

func createQuiz() Quiz {
	problems := sampleProblems()
	return New(problems)
}

func runQuiz(buffer io.Writer, quiz *Quiz) {
	answers := strings.NewReader("10\n2\n")
	quiz.Run(buffer, answers)
}
