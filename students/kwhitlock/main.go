package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var totalScore = 0
var numberQuestionsAsked = 0

func getUserAnswer() (text string) {
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return answer
}

func sendQuestion(text string) (answer string) {
	questionAnswer := strings.Split(text, ",")
	fmt.Println(questionAnswer[0])
	numberQuestionsAsked++
	return questionAnswer[1]
}

func compareAnswers(answer, userAnswer string) {
	if answer == strings.TrimRight(userAnswer, "\n") {
		fmt.Println("Correct!")
		totalScore++
	} else {
		fmt.Println("Incorrect :o(")
	}
}

func timer(wg *sync.WaitGroup, lengthOfTime int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(lengthOfTime))
}

func quiz(wg *sync.WaitGroup, scanner *bufio.Scanner) {
	defer wg.Done()
	for scanner.Scan() {
		answer := sendQuestion(scanner.Text())
		userAnswer := getUserAnswer()
		compareAnswers(answer, userAnswer)
	}
}

func main() {
	questions := flag.String("questions", "./problems.csv", "Filepath/name of CSV file where the questions are stored.")
	lengthOfTime := flag.Int("time", 30, "Time in seconds to run quiz for.")

	flag.Parse()

	fmt.Println("Running Quiz")

	file, err := os.Open(*questions)
	if err != nil {
		fmt.Println("Could not open file")
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup

	wg.Add(1)
	fmt.Println("Ready to start quiz? (press enter)")
	getUserAnswer()
	go timer(&wg, *lengthOfTime)
	go quiz(&wg, scanner)
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}
	fmt.Printf("**End of quiz**\nYou scored: %v/%v\n", totalScore, numberQuestionsAsked)

}
