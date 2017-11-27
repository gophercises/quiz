package main

import (
	"bufio"
	"os"

	"github.com/quiz/solution"
)

func main() {
	quiz := solution.MyQuiz{
		File:   "problems.csv",
		Reader: bufio.NewReader(os.Stdin),
	}

	quiz.Start()
}
