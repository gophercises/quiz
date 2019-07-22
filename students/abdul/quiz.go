package quiz

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

//Number of Questions to ask
const totalQuestions = 5

//Question struct that stores question with answer
type Question struct {
	question string
	answer   string
}

func main() {
	filename, timeLimit := readArguments()
	f, err := openFile(filename)
	if err != nil {
		return
	}
	questions, err := readCSV(f)

	if err != nil {
		// err := fmt.Errorf("Error in Reading Questions")
		fmt.Println(err.Error())
		return
	}

	if questions == nil {
		return
	}
	score, err := askQuestion(questions, timeLimit)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Your Score %d/%d\n", score, totalQuestions)
}

func readArguments() (string, int) {
	filename := flag.String("filename", "problem.csv", "CSV File that conatins quiz questions")
	timeLimit := flag.Int("limit", 30, "Time Limit for each question")
	flag.Parse()
	return *filename, *timeLimit
}

func readCSV(f io.Reader) ([]Question, error) {
	// defer f.Close() // this needs to be after the err check
	allQuestions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	numOfQues := len(allQuestions)
	if numOfQues == 0 {
		return nil, fmt.Errorf("No Question in file")
	}

	var data []Question
	for _, line := range allQuestions {
		ques := Question{}
		ques.question = line[0]
		ques.answer = line[1]
		data = append(data, ques)
	}

	return data, nil
}

func openFile(filename string) (io.Reader, error) {
	return os.Open(filename)
}
func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input <- result
	}
}

func askQuestion(questions []Question, timeLimit int) (int, error) {
	totalScore := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	done := make(chan string)

	go getInput(done)

	for i := range [totalQuestions]int{} {
		ans, err := eachQuestion(questions[i].question, questions[i].answer, timer.C, done)
		if err != nil && ans == -1 {
			return totalScore, nil
		}
		totalScore += ans

	}
	return totalScore, nil
}

func eachQuestion(Quest string, answer string, timer <-chan time.Time, done <-chan string) (int, error) {
	fmt.Printf("%s: ", Quest)

	for {
		select {
		case <-timer:
			return -1, fmt.Errorf("Time out")
		case ans := <-done:
			score := 0
			if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), answer) == 0 {
				score = 1
			} else {
				return 0, fmt.Errorf("Wrong Answer")
			}

			return score, nil
		}
	}
}
