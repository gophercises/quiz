package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	timeOut  int
	quizFile string
	shuffle  bool
)

type Quiz struct {
	Question string
	Expected string
	Response string
}

func init() {
	flag.IntVar(&timeOut, "limit", 30, "Timeout for the quiz seconds")
	flag.StringVar(&quizFile, "csv", "problems.csv", "The problem file")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the questions")
	flag.Parse()
}

func main() {
	quizFd, err := os.Open(quizFile)
	stopCh := make(chan bool, 1)
	resultCh := make(chan int, 1)
	if err != nil {
		log.Fatal(err)
	}
	quizes, err := readCSV(quizFd)
	if err != nil {
		log.Fatal("Unable to parse csv file")
	}
	duration := time.Duration(timeOut) * time.Second

	fmt.Printf("Ready for quiz, press a enter to start")
	var key string
	fmt.Scanln(&key)
	timer := time.NewTimer(duration)
	if shuffle {
		shuffleQuiz(quizes)
	}
	go runQuiz(quizes, stopCh, resultCh)
	correct := 0
	for {
		select {
		case <-timer.C:
			fmt.Println("\nTimer buzzz")
			stopCh <- true
		case val, more := <-resultCh:
			if more {
				correct = val
			} else {
				showQuizResults(quizes, correct)
				return
			}
		}
	}
}

func readCSV(quizFd *os.File) ([]Quiz, error) {
	r := csv.NewReader(quizFd)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	quizes := make([]Quiz, 0, len(records))
	for _, a := range records {
		quizes = append(quizes, Quiz{a[0], strings.TrimSpace(a[1]), ""})
	}
	return quizes, nil
}

func runQuiz(quizes []Quiz, stopCh <-chan bool, resultCh chan<- int) {
	correct := 0
	go func() {
		for i, _ := range quizes {
			fmt.Printf("%s = ", quizes[i].Question)
			_, err := fmt.Fscanln(os.Stdin, &quizes[i].Response)
			if err == io.EOF {
				break
			}
			if quizes[i].Response == quizes[i].Expected {
				correct++
			}
			resultCh <- correct
		}
		close(resultCh)
	}()
	<-stopCh
	close(resultCh)
}

func shuffleQuiz(quizes []Quiz) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizes), func(i, j int) {
		quizes[i], quizes[j] = quizes[j], quizes[i]
	})
}

func showQuizResults(quiz []Quiz, correct int) {
	fmt.Printf("\nYou scored %d of %d\n", correct, len(quiz))
	for _, q := range quiz {
		fmtStr := "%s = %s %c\n"
		if q.Response != q.Expected {
			fmt.Printf(fmtStr, q.Question, q.Response, '\u2717')
		} else {
			fmt.Printf(fmtStr, q.Question, q.Response, '\u2713')
		}
	}
}
