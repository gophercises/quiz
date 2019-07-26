package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"flag"
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

func main(){
	questions := flag.String("questions", "./problems.csv", "Filepath/name of CSV file where the questions are stored.")
	flag.Parse()

	fmt.Println("Running Quiz")

	file, err := os.Open(*questions)
	if err != nil {
		fmt.Println("Could not open file")
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		answer := sendQuestion(scanner.Text())
		userAnswer := getUserAnswer()
		compareAnswers(answer, userAnswer)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}
	fmt.Printf("**End of quiz**\nYou scored: %v/%v\n", totalScore, numberQuestionsAsked)

}