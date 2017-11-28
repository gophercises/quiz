package business

import (
	"bufio"
	"fmt"

	"github.com/quiz/reader"
)

type MyQuiz struct {
	File  string "path of the File"
	timer int    "time in seconds"
	name  string "quiz name"

	Reader     *bufio.Reader "Reader for reading user answers"
	tCorrect   int
	tIncorrect int
	tQuestions int
}

func (quiz *MyQuiz) run(questions map[string]string) {
	if quiz.name != "" {
		fmt.Printf("Starting %v quiz", quiz.name)
	}
	quiz.tQuestions = len(questions)
	for ques, ans := range questions {
		display(ques)
		userAnswer := getUserAns(quiz.Reader)
		correct, incorrect := validate(ans, userAnswer)
		quiz.tCorrect += correct
		quiz.tIncorrect += incorrect
	}
}

func validate(one string, two string) (correct int, incorrect int) {
	if one == two {
		return 1, 0
	}
	return 0, 1
}
func getUserAns(reader *bufio.Reader) string {
	ans, ok := reader.ReadString('\n')
	if ok != nil {
		fmt.Errorf("Error while reading user answer")
	}
	return ans
}
func display(question string) {
	fmt.Println(question)
}

func (quiz *MyQuiz) Start(over chan bool) {
	dat := reader.ReadCsv(quiz.File)
	quiz.run(dat)
	over <- true
}
func (quiz *MyQuiz) Result() {
	fmt.Printf("Correct: %v, incorrect: %v, total: %v\n", quiz.tCorrect, quiz.tIncorrect, quiz.tQuestions)
}

type Quiz interface {
	start()
	result()
}
