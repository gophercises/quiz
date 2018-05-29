package quiz

import (
	"fmt"
	"time"

	"github.com/fedepaol/quiz/interaction"
)

// Question represents a single quiz question with answer.
type Question struct {
	Question string
	Answer   string
}

// QuestionService implements all the methods related to a single quiz question.
type QuestionService interface {
	Ask() bool
}

// Quiz represents a quiz run with all the questions and answers.
type Quiz struct {
	questions   []Question
	Results     []bool
	asker       interaction.Asker
	goodreplies int
}

// QuizService holds all the methods that can be applied to a Quiz.
type QuizService interface {
	Run()
	AddQuestion(Question)
}

// Run runs an iteration of a quiz.
func (q *Quiz) Run() {
	replies := make(chan bool)

	go func() {
		for _, qq := range q.questions {
			reply := q.asker.Ask(qq.Question)

			if reply == qq.Answer {
				replies <- true
			} else {
				replies <- false
			}
		}
	}()

	timer := time.NewTimer(2 * time.Second)

T:
	for {
		select {
		case success := <-replies:
			if success {
				q.goodreplies++
			}
		case <-timer.C:
			break T
		}
	}
	msg := fmt.Sprintf("You got %d out of %d right questions", q.goodreplies, len(q.questions))
	q.asker.Notify(msg)
	q.goodreplies = 0
}

// AddQuestion adds a question to the quiz.
func (q *Quiz) AddQuestion(question Question) {
	q.questions = append(q.questions, question)
}
