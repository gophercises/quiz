package quiz

import (
	"fmt"
	"io"

	"github.com/gophercises/quiz/students/hackeryarn/problem"
)

// Quiz represents the quiz to be given to the user
type Quiz struct {
	problems     []problem.Problem
	rightAnswers int
}

// Run runs the quiz for all the problems keeping track of correct answers
func (q *Quiz) Run(w io.Writer, r io.Reader) {
	for _, problem := range q.problems {
		problem.AskQuestion(w)
		correct := problem.CheckAnswer(r)
		if correct {
			q.rightAnswers++
		}
	}

	q.PrintResults(w)
}

// PrintResults outputs the results of the quiz
func (q Quiz) PrintResults(w io.Writer) {
	fmt.Fprintf(w, "You got %d questions right!\n", q.rightAnswers)
}

// New creates a new quiz from the supplied slice of problems
func New(problems []problem.Problem) Quiz {
	return Quiz{
		problems:     problems,
		rightAnswers: 0,
	}
}
