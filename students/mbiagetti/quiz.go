package main

import (
	"flag"
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"time"
	"fmt"
)

type Question struct{
	question string
	answer string
}

type Quiz struct {
	questions []Question
}

func main() {
	var correctAnswer int
	var quiz Quiz

	fileName := flag.String("csv", "problems.csv", `a csv file in format "question,answer"`)
	limit := flag.Int("limit", 30, "the time limit for the quiz")

	flag.Parse()

	fmt.Println("Welcome to the quiz")
	fmt.Printf("Reading question from %s. Time limit: %d\n", *fileName, *limit)
	fmt.Println("Press enter to start the game!")

	csvFile, _ := os.Open(*fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		quiz.questions = append(quiz.questions, Question{
			question: line[0],
			answer:  line[1],
		})
	}

	doneChan := make(chan bool)

	fmt.Scanln()

	// timeout func
	go func() {
		time.Sleep(time.Second * time.Duration(*limit))
		fmt.Println("Timeout!")
		doneChan <- true
	}()

	// quiz func
	go func() {
		for _,k := range quiz.questions {
			var text string
			fmt.Printf("%s: ", k.question)
			fmt.Scanln(&text)
			if k.answer == text {
				correctAnswer++
			}
		}
		doneChan <- true
	}()

	for {
		select {
		case <- doneChan:
			fmt.Println("Done")
			fmt.Printf("Correct answer: %d\n", correctAnswer)
			return
		}
	}
}
