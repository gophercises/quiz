package problem

import (
	"strings"
)

// Question represents the question/problem set for the player of the quiz to answer.
// It contains both question text and the correct answer that answers the question.
type Question struct {
	questionText      string
	correctAnswerText string
}

// NewQuestion creates a new quiz question by providing the question string and
// correctAnswer that answers the question
func NewQuestion(question string, correctAnswer string) *Question {
	q := Question{question, correctAnswer}
	return &q
}

// Question retrieves the question text
func (q *Question) Question() string {
	return q.questionText
}

// Give attempts to answer the question by providing the proposed answer.
// Returns boolean whether the answer was correct or not
func (q *Question) Give(answer string) bool {
	// ignore case when comparing answers
	if strings.ToLower(answer) == strings.ToLower(q.correctAnswerText) {
		return true
	}

	return false
}
