package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	term "github.com/nsf/termbox-go"

	"os"
	"time"
)

var (
	score           int
	questionsNumber int
	filename        *string
	duration        *int
	Answers         []string
	userQuestions   []string
)

func startTheTimer() {
	for alive := true; alive; {
		timer := time.NewTimer(time.Duration(*duration) * time.Second)
		select {
		case <-timer.C:
			alive = false
			fmt.Println("\nTime is up .")
			fmt.Println("you got", score, "out of", questionsNumber)
			os.Exit(0)

		}
	}
}

func lineCounter(r *csv.Reader) (int, error) {
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		questionsNumber++
	}
	return questionsNumber, nil
}

func StartTheQuiz(filename *string) {

	//read csv file
	file, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(0)
	}

	//read from  file
	r := csv.NewReader(strings.NewReader(string(file)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		userQuestions = append(userQuestions, record[0])
		Answers = append(Answers, record[1])

	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press 1 to start the quiz and 2 to exit")
	userAnswer, err := reader.ReadString('\n')
	if strings.TrimRight(userAnswer, "\n") == "1" {
		go startTheTimer()
		for i := 0; i < len(userQuestions); i++ {
			fmt.Printf("What is answer for this %s :  ", userQuestions[i])
			userAnswer, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			if strings.TrimRight(userAnswer, "\n") == Answers[i] {
				score = score + 1
			}
		}
	} else {
		os.Exit(0)
	}

	fmt.Println("you got", score, "out of", len(userQuestions))

}

func reset() {
	term.Sync() // cosmestic purpose
}
func main() {

	filename = flag.String("filename", "problems.csv", "file name of the problems,only csv file.questions/answers format")
	duration = flag.Int("duration", 30, "duration of quiz in seconds ")
	flag.Parse() //parse flags from command line.

	StartTheQuiz(filename)
}
