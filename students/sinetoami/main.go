// Package main provides a command that starts a quiz based on a CSV file.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

// Problem represents a quis' problem structure.
type Problem struct {
	question, answer string
}

// QuizProblems stores each problem from a given CSV file to a variable.
// If the shuffle flag is true will randomize the problems order.
func QuizProblems(lines [][]string, shuffle bool) []Problem {
	problems := make([]Problem, len(lines))
	rand.Seed(time.Now().UTC().UnixNano())

	// create a list with random items (int)
	perm := rand.Perm(len(lines))

	if !shuffle {
		sort.Ints(perm[:])
	}

	for i, v := range perm {
		problems[v] = Problem{lines[i][0], lines[i][1]}
	}

	return problems
}

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds (default 30)")
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz order each time it is run if pass value 'true' (default false)")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Errorf("filed to read lines from file: %v", err)
	}

	score := 0
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	problems := QuizProblems(lines, *shuffle)

quizloop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)

		// if the user don't type the answer before the timer is done
		// will terminate the quiz.
		answerChan := make(chan string)
		go func() {
			userAnswer := ""
			fmt.Scanf("%s\n", &userAnswer)
			answerChan <- strings.TrimSpace(userAnswer)
		}()

		select {
		case an := <-answerChan:
			if an == problem.answer {
				score++
			}
		case <-timer.C:
			break quizloop
		}
	}

	fmt.Printf("You scored %d out of %d.", score, len(problems))
}
