package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type problem struct {
	question string
	answer string
}

var corrects int
var failed int

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeDuration := flag.Int("duration", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	_ = timeDuration
	
	file, err := os.Open(*csvFileName)
	defer file.Close()
	if  err != nil {
		exit(fmt.Sprintf("Error while opening csv file: %s\n", *csvFileName))
	}

	fileContent := csv.NewReader(file)
	lines, err := fileContent.ReadAll()
	if err != nil {
		exit(err.Error())
	}
	problems := parseLines(lines)

	for i, problem := range problems {
		question := problem.question
		answer, _ := strconv.Atoi(problem.answer)
		fmt.Printf("Problem %d: %s = ", i+1, question)
		
		scanner  := bufio.NewScanner(os.Stdin);
		scanner.Scan();
		input , err := strconv.Atoi(scanner.Text());
		if err != nil {
			fmt.Println("\nAn error ocurred while trying to pass answer.");
			return;
		}

		if answer == input {
			corrects++
		} else {
			failed++
		}
	}
	fmt.Println("Correct answers:", corrects)
	fmt.Println("Incorrect answers:", failed)
}

func parseLines(problems [][]string) []problem {
	ret := make([]problem, len(problems))
	for i, line := range problems {
		ret[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}