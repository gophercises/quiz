package main

import (
	"fmt"
	"encoding/csv"
	"strings"
	"io/ioutil"
	"log"
	"bufio"
	"os"
	"time"
)

type QuestionSet struct {
	questions []string
	actualAnswers []string
	userAnswers []string
}

func fetchResult(qsf QuestionSet, c chan int) {
	for i, _ := range qsf.userAnswers {
		if strings.Compare(qsf.userAnswers[i], qsf.actualAnswers[i]) == 1 {
			c <- 1;
		}
		c <- 0;
	}
	close(c)
}

func main() {
	var timeLimit int = 10
	// Reads as []uint8
	data, err := ioutil.ReadFile("problems.csv")
	if (err != nil) {
		fmt.Println(err)
	}
	// Parses the csv file into *csv.Reader
	r := csv.NewReader(strings.NewReader(string(data)))
	// Read from the reader
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// Store in our struct
	qs := new(QuestionSet)
	for _, record := range records {
		qs.questions = append(qs.questions, record[0])
		qs.actualAnswers = append(qs.actualAnswers, record[1])
	}
	
	fmt.Println("Parsed")
	// Store answers from the user, timer starts
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please answer the following questions\n")
	timer := time.NewTimer(time.Duration(timeLimit)*time.Second)
	
	// timer returns current time after some duration
	loop: for _, question := range qs.questions {
		fmt.Println(question);
		answer, _ := inputReader.ReadString('\n')
		select {
		case <-timer.C:
			fmt.Println("Times up!!!")
			break loop
		default:
			qs.userAnswers = append(qs.userAnswers, answer) 
		}
	}

	// run go routines to fetch result
	ch := make(chan int, 10)
	go fetchResult(*qs, ch)
	res := 0
	for i := range ch {
		res = res + i
	}
	fmt.Println("Your score is: ", res)
}