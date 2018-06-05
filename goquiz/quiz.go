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
	questions    []Question
	Asker        interaction.Asker
	questionsnum int
}

type Result struct {
	goodreplies    int
	questionsasked int
}

// QuizService holds all the methods that can be applied to a Quiz.
type QuizService interface {
	Run()
	AddQuestion(Question)
}

// Run runs an iteration of a quiz.
func (q *Quiz) Run(timeout <-chan time.Time) (res Result) {
	replies := make(chan bool)

	go func() {
		for _, qq := range q.questions {
			reply := q.Asker.Ask(qq.Question)

			if reply == qq.Answer {
				replies <- true
			} else {
				replies <- false
			}
		}
		close(replies)
	}()

T:
	for {
		select {
		case success, ok := <-replies:
			if !ok {
				break T
			}

			if success {
				res.goodreplies++
			}
			res.questionsasked++
		case <-timeout:
			break T
		}
	}
	return
}

// AddQuestion adds a question to the quiz.
func (q *Quiz) AddQuestion(question Question) {
	q.questions = append(q.questions, question)
}

func (r Result) String() (res string) {
	res = fmt.Sprintf("%d good replies out of %d questions asked", r.goodreplies, r.questionsasked)
	return
}
