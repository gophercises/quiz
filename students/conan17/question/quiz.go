package question

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

var _ Quiz = &quiz{}
var _ Quizzes = &quizzes{}

type quiz struct {
	problem string
	answer  string
	reader  io.Reader
	write   io.Writer
}

func (problem *quiz) Ask() {
	fmt.Fprintf(problem.write, "Quesion: %s\n", problem.problem)
}

func (problem *quiz) Answer() bool {
	scanner := bufio.NewScanner(problem.reader)
	scanner.Scan()
	respond := scanner.Text()
	if strings.EqualFold(strings.TrimSpace(respond), problem.answer) {
		return true
	}
	return false
}

func NewQuiz(problem string, answer string) Quiz {
	return &quiz{
		problem: problem,
		answer:  answer,
		reader:  os.Stdin,
		write:   os.Stdout,
	}
}

type quizzes struct {
	quizzes  []Quiz
	current  int
	answered map[int]bool
	writer   io.Writer
	random   bool
}

func (problem *quizzes) Random(r bool) {
	problem.random = r
}

func (problem *quizzes) Summary() {
	rightNum := 0
	totalNum := 0
	for _, ok := range problem.answered {
		if ok {
			rightNum++
		}
		totalNum++
	}
	fmt.Fprintf(problem.writer, "Answered %d/%d questions.\n", rightNum, totalNum)
}

func (problem *quizzes) Launch() {
	if problem.current >= len(problem.quizzes) || problem.current < 0 {
		return
	}
	if problem.answered == nil {
		problem.answered = make(map[int]bool)
	}
	quiz := problem.quizzes[problem.current]
	fmt.Fprintf(problem.writer, "%d. ", problem.Size()-problem.Remain())
	quiz.Ask()
	problem.answered[problem.current] = quiz.Answer()

}

func (problem *quizzes) Size() int {
	return len(problem.quizzes)
}

func (problem *quizzes) Remain() int {
	return problem.Size() - len(problem.answered)
}

func (problem *quizzes) Next() error {
	if problem.answered == nil {
		problem.answered = make(map[int]bool)
	}
	if problem.Remain() <= 0 {
		return io.EOF
	}

	if problem.random {
		rand.Seed(time.Now().Unix())
		problem.current = rand.Intn(len(problem.quizzes))
		for _, ok := problem.answered[problem.current]; ok; _, ok = problem.answered[problem.current] {
			problem.current = rand.Intn(len(problem.quizzes))
		}
	} else {
		if problem.current >= len(problem.quizzes) {
			problem.current = 0
		}
		for _, ok := problem.answered[problem.current]; ok; _, ok = problem.answered[problem.current] {
			problem.current++
		}
	}
	problem.answered[problem.current] = false
	return nil
}

func NewQuizzes(q ...Quiz) Quizzes {
	return &quizzes{
		writer:   os.Stdout,
		quizzes:  q,
		answered: make(map[int]bool),
	}
}
