package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/fedepaol/quiz/interaction"
	"github.com/fedepaol/quiz/parser"
)

func main() {
	csvname := flag.String("file", "problems.csv", "the name of the csv file to be used")
	timeout := flag.Int("timeout", 30, "the max duration of the quiz in seconds")

	flag.Parse()
	quiz, err := parser.ParseFile(*csvname)
	if err != nil {
		fmt.Errorf("error %v", err)
		return
	}
	quiz.Asker = interaction.NewAsker()

	to := time.NewTimer(time.Duration(*timeout) * time.Second)
	quiz.Run(to.C)
}
