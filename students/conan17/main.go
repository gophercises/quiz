package main

import (
	"flag"
	"time"

	"./question"
)

var (
	isRandom     = flag.Bool("rand", false, "randomize questions")
	csvPath      = flag.String("path", "./problem.csv", "Path to csv file containing quiz.")
	timeLimitStr = flag.String("limit", "30s", "Set the total time allowed for the quiz.")
)

func main() {
	var (
		timeLimit time.Duration
		err       error
		quizzes   question.Quizzes
		syncChan  = make(chan struct{}, 0)
	)
	flag.Parse()
	if timeLimit, err = time.ParseDuration(*timeLimitStr); err != nil {
		panic(err)
	}
	quizzes, err = question.CSVQuizzes(*csvPath)
	if err != nil {
		panic(err)
	}
	quizzes.Random(*isRandom)
AnswerBegin:
	for {
		if err = quizzes.Next(); err != nil {
			break
		}
		go func() {
			quizzes.Launch()
			<-syncChan
		}()
		select {
		case <-time.After(timeLimit):
			break AnswerBegin
		case syncChan <- struct{}{}:
		}
	}
	quizzes.Summary()
}
