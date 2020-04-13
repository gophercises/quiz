package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type qna struct {
	q string
	a string
}

func main() {

	fmt.Println("Starting Quiz")
	for i := 0 ; i<10; i++ {
		time.Sleep(250*time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()

 	csvFileName := flag.String("csv", "Problems.csv", "csv file in the formate of questions and answer")
	timeLimit := flag.Int("limit", 30, "this provides the time limit for the question to solve")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(err)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(err)
	}

	timer := time.NewTimer(time.Duration(*timeLimit)*time.Second)
	
	quiz := make([]qna, len(lines))

	for i , line:= range lines {
		quiz[i].q = line[0]
		quiz[i].a = strings.TrimSpace(line[1])
		quiz[i].a = strings.TrimLeft(quiz[i].a,"0")
	}

	var count int
	var questionNumber int

	programLoop:
	for i := range lines {
		fmt.Printf("Question %v: %s = ", i+1, quiz[i].q)
		questionNumber++
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanln(&answer)
			answer = strings.TrimSpace(answer)
			answer = strings.TrimLeft(answer,"0")
			answer = strings.ToUpper(answer)
			fmt.Printf("Your answer to the question is %v\n",answer)
			//fmt.Printf("expected answer was %v\n",quiz[i].a)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYour score is %v out of %v\n",count, len(lines))
			break programLoop
		case answer := <-answerCh:
			if questionNumber == len(lines) {
				fmt.Printf("\nYour score is %v out of %v\n",count, len(lines))
				break programLoop
			}
			if answer == quiz[i].a {
				//fmt.Println("correct")
				count++
			}
		}
	}
}

func exit(msg error) {
	fmt.Println(msg)
	os.Exit(1)
}
