package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	welcomeMsg()
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Cannot Open the file.")
	}
	fmt.Println("File Opened sucessfully.")
	defer csvFile.Close()

	readLine := csv.NewReader(csvFile)
	lines, err := readLine.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	questions := lineParser(lines)

	correct := 0
	for i, question := range questions {
		fmt.Printf("Question number %d: %s = ", i+1, question.Ques)
		var response string
		fmt.Scanf("%s\n", &response)
		if response == question.Ans {
			correct += 1
		}

	}

	fmt.Printf("You attempted %d out of %d questions correctly. \n", correct, len(questions))

}

// A welcome Message
func welcomeMsg() {
	welcome := `
	-------------------------------
	|                             |
	|  Welcome to the Math Quiz   |
	!                             |
	-------------------------------`
	fmt.Println(welcome)
}

type problem struct {
	Ques string
	Ans  string
}

// Parsing each lines of the csv file
func lineParser(lines [][]string) []problem {
	prob := make([]problem, len(lines))
	for i, line := range lines {
		prob[i] = problem{
			Ques: line[0],
			Ans:  line[1],
		}
	}
	return prob
}
