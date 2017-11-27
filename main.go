package main

import (
	"bufio"
	"os"

	"flag"

	"github.com/quiz/solution"
)

func main() {
	file := flag.String("file", "problems.csv", "file containing quiz questions")
	if *file == "problems.csv" {
		file = flag.String("f", "problems.csv", "file containing quiz questions")
	}
	flag.Parse()

	quiz := solution.MyQuiz{
		File:   *file,
		Reader: bufio.NewReader(os.Stdin),
	}

	quiz.Start()
}
