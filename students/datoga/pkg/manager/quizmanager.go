package manager

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/datoga/quiz/pkg/model"
)

type QuizManager struct {
	quiz  model.Quiz
	chEnd chan bool
	step  int
	ok    int
}

func NewQuizManager(quiz model.Quiz) *QuizManager {
	chEnd := make(chan bool, 1)

	return &QuizManager{
		quiz:  quiz,
		chEnd: chEnd,
	}
}

func (manager *QuizManager) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(manager.quiz), func(i, j int) {
		manager.quiz[i], manager.quiz[j] = manager.quiz[j], manager.quiz[i]
	})
}

func (manager *QuizManager) Quiz() {
	for i, v := range manager.quiz {
		fmt.Printf("Question %d: %s: ", (i + 1), v.Question)

		answer := ""

		_, err := fmt.Scanln(&answer)

		if err != nil {
			fmt.Println(err)
			continue
		}

		answer = clean(answer)

		manager.step++

		if answer == v.Solution {
			manager.ok++
		}
	}

	manager.chEnd <- true
}

func clean(answer string) string {
	answer = strings.Trim(answer, " \t\n\r")

	answer = strings.ToLower(answer)

	return answer
}

func (manager QuizManager) GetResults() (step int, questions int, ok int) {
	return manager.step, len(manager.quiz), manager.ok
}

func (manager QuizManager) NotifyEnd() chan bool {
	return manager.chEnd
}
