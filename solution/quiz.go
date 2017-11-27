package solution

import (
	"bufio"
	"fmt"
)

type MyQuiz struct {
	File  string "path of the File"
	timer int    "time in seconds"
	name  string "quiz name"

	Reader     *bufio.Reader "Reader for reading user answers"
	tCorrect   int
	tIncorrect int
}

func (quiz *MyQuiz) run(questions map[string]string) {
	if quiz.name != "" {
		fmt.Printf("Starting %v quiz", quiz.name)
	}

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

func (quiz *MyQuiz) Start() {
	dat := ReadCsv(quiz.File)
	quiz.run(dat)
	quiz.result()
}
func (quiz *MyQuiz) result() {
	fmt.Printf("Correct: %v, incorrect: %v, total: %v\n", quiz.tCorrect, quiz.tIncorrect, quiz.tCorrect+quiz.tIncorrect)
}

type Quiz interface {
	start()
	result()
}
