package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// QuestionPair includes both the question and the answer values
type QuestionPair struct {
	question string
	answer   string
}

func main() {
	filePrt := flag.String("file", "problems.csv", "is string")
	tmrPrt := flag.String("timer", "30", "is number")
	flag.Parse()
	qp, _ := readQuestions(*filePrt)

	r := bufio.NewReader(os.Stdin)
	ca := 0

	fmt.Println("Press Enter to start. the timer total is set at", *tmrPrt)
	r.ReadString('\n')
	t, _ := strconv.Atoi(*tmrPrt)

	delay := time.Duration(t) * time.Second
	compeleted := false

	tmr := time.NewTimer(delay)

qLoop:
	for i := 0; i < len(qp); i++ {
		fmt.Println(qp[i].question)
		fmt.Print("-> ")

		ansChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansChan <- answer
		}()

		select {
		case <-tmr.C:
			break qLoop
		case answer := <-ansChan:
			if answer == qp[i].answer {
				ca++
			}
			if i == len(qp)-1 {
				compeleted = true
			}
		}
	}

	if !compeleted {
		fmt.Println("You ran out of time!")
	}

	fmt.Println(ca, "out of", len(qp), "correct answers")
}

// readQuestions reads a csv file and returns a slice of QuestionPair
func readQuestions(fn string) ([]QuestionPair, error) {
	f, err := os.Open(fn)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	r := csv.NewReader(f)
	ss, err := r.ReadAll()

	qp := make([]QuestionPair, len(ss))

	for i, s := range ss {
		qp[i] = QuestionPair{s[0], s[1]}
	}

	return qp, nil
}
