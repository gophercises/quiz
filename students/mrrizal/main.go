package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type problem struct {
	question string
	answer   string
}

func parseCSVFile(csvLines [][]string) []problem {
	result := make([]problem, len(csvLines))
	for i, line := range csvLines {
		result[i].question = strings.TrimSpace(line[0])
		result[i].answer = strings.TrimSpace(line[1])
	}
	return result
}

func readCSVFile(filename string) [][]string {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	csvLines, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	return csvLines
}

func quiz(problems []problem, limit int) {
	reader := bufio.NewReader(os.Stdin)
	correctAsnwer := 0
	timer := time.NewTimer(time.Duration(limit) * time.Second)

quizloop:
	for counter, problem := range problems {
		fmt.Printf("%d. %s = ", counter+1, problem.question)
		answerChannel := make(chan string)
		go func() {
			userAnswer, _ := reader.ReadString('\n')
			answerChannel <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTimeout ...")
			break quizloop
		case userAnswer := <-answerChannel:
			if strings.TrimSpace(userAnswer) == problem.answer {
				correctAsnwer++
			}
		}
	}

	fmt.Printf("\ncorrect answer %d from %d\n", correctAsnwer, len(problems))
}

func main() {
	var filename string
	var limit int
	flag.StringVar(&filename, "filename",
		"problems.csv", "csv file with format `question,answer")
	flag.IntVar(&limit, "limit", 30, "limit time in second")
	flag.Parse()

	csvLines := readCSVFile(filename)

	// randomize
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(csvLines), func(i, j int) {
		csvLines[i], csvLines[j] = csvLines[j], csvLines[i]
	})

	problems := parseCSVFile(csvLines)
	quiz(problems, limit)
}
