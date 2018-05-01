package main

import (
	"flag"
	"fmt"
	"quiz/game"
	"quiz/timer"
	"time"
)

const csvFlagDescription = `a csv file in the format of 'question,answer'`
const limitFlagDescription = `the time limit for the quiz in seconds`
const shuffleFlagDescription = `provide quiz questions in random order`

func main() {
	// parse csv filename tag possibly given as a shell command switch
	csvFile := flag.String("csv", "problems.csv", csvFlagDescription)
	timeoutSeconds := flag.Int("limit", 30, limitFlagDescription)
	shuffleQuiz := flag.Bool("shuffle", false, shuffleFlagDescription)
	flag.Parse()

	engine := game.NewEngine()
	engine.SetQuizRandomness(*shuffleQuiz)
	err := engine.SetQuizProblems(*csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create a new runner to time the StartGame() job
	runner := timer.NewRunner(time.Duration(*timeoutSeconds) * time.Second)
	runner.Add(engine.StartGame)
	err = runner.Run() // will run engine.StartGame()
	if err != nil {
		fmt.Println(err)
	}
	engine.PrintGameResults()

	return
}
