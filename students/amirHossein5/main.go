package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type question struct {
	question string
	answer   string
}

func main() {
	flags := initFlags()
	questions := parseCSVQuestions(flags["csvFileName"].(string), flags["shuffle"].(bool))
	correctCount := 0

questionsLoop:
	for i, question := range questions {
		answerC := make(chan string)
		timer := time.NewTimer(time.Duration(flags["timeLimit"].(int)) * time.Second)

		fmt.Printf("#%d- %s ", i+1, question.question)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerC <- answer
		}()

		select {
		case answer := <-answerC:
			if answer == question.answer {
				correctCount++
			}
			timer.Stop()
		case <-timer.C:
			fmt.Println("reached time limit!")
			break questionsLoop
		}
	}

	fmt.Printf("\n%d out of %d \n", correctCount, len(questions))
}

func initFlags() map[string]any {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of question, answer")
	shuffle := flag.Bool("shuffle", true, "shuffle sort of questions each time")
	timeLimit := flag.Int("time-limit", 20, "time limit for the quiz in seconds")

	flag.Parse()

	return map[string]any{
		"csvFileName": *csvFileName,
		"shuffle":     *shuffle,
		"timeLimit":   *timeLimit,
	}
}

func parseCSVQuestions(filename string, shuffle bool) []question {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to open the CSV file: %s", err)
	}

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		log.Fatalf("Failed to parse CSV: %s", err)
	}

	if shuffle {
		shuffleCSVLines(&lines)
	}

	return getQuestionsFromLines(lines)
}

func shuffleCSVLines(lines *[][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*lines), func(i, j int) {
		(*lines)[i], (*lines)[j] = (*lines)[j], (*lines)[i]
	})
}

func getQuestionsFromLines(lines [][]string) []question {
	questions := make([]question, len(lines))

	for i, line := range lines {
		questions[i] = question{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return questions
}
