package main

import (
	"bufio"
	"os"

	"flag"

	"time"

	"fmt"

	"github.com/quiz/business"
)

func main() {

	file := flag.String("file", "problems.csv", "file containing quiz questions")
	if *file == "problems.csv" {
		file = flag.String("f", "problems.csv", "file containing quiz questions")
	}

	limit := flag.Int("limit", 10, "max quiz time limit")
	if *limit == 10 {
		limit = flag.Int("l", 10, "max quiz time limit")
	}
	flag.Parse()

	quiz := business.MyQuiz{
		File:   *file,
		Reader: bufio.NewReader(os.Stdin),
	}

	var tm, qz chan bool
	tm = make(chan bool)
	qz = make(chan bool)
	go quiz.Start(qz)
	go startTimer(*limit, tm)
	select {
	case <-tm:
		fmt.Println("Your time is over.")

	case <-qz:
		fmt.Println("All questions are answered.")
	}
	quiz.Result()
}
func startTimer(tm int, over chan bool) {
	time.Sleep(time.Duration(tm) * time.Second)
	over <- true
}
