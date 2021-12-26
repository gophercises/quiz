package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("filename", "problems.csv", "file containing the set of problems")
	timeLimit := flag.Int("limit", 30, "quiz time limit")
	flag.Parse()

	quizzes := getQuiz(readFile(*filename))
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0
	for _, quiz := range quizzes {
		fmt.Print(quiz.question + " : ")
		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			printResults(score, len(quizzes))
			return
		case ans := <-ansCh:
			if strings.TrimSpace(quiz.answer) == strings.TrimSpace(ans) {
				score++
			}
		}
	}
	printResults(score, len(quizzes))
}
func printResults(score, total int) {
	fmt.Printf("You scored %d out of %d\n", score, total)
}

func getQuiz(records [][]string) []Quiz {
	quizzes := make([]Quiz, 0)
	for _, record := range records {
		quizzes = append(quizzes, Quiz{question: record[0], answer: record[1]})
	}
	return quizzes
}

func readFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV "+filePath, err)
	}
	return records
}
