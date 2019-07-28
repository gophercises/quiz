package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// Grab the problems file from the cli
	file := flag.String("filename", "problems.csv", "Pass in the filename of the csv file.")
	// quizTime := flag.Int("time", 30, "This represents how much time a user will have to complete the quiz.")
	var rightAnswers int
	var wrongAnswers int
	var total int
	flag.Parse()

	// Open the csv file
	inputBytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Print(err)
	}

	// Convert the problems file from bytes to a string
	data := string(inputBytes)

	// First, we create a Reader type out of the data (string) and then we pass it to the csv.NewReader()
	// so that we can read the data from the csv file.
	r := csv.NewReader(strings.NewReader(data))
	r.Comma = ','

	// Read all of the records into memory
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// This reads in inputs from the STDIN device to the NewReader memory
	reader := bufio.NewReader(os.Stdin)

	start := time.Now()
	for _, record := range records {
		fmt.Printf("What is %s?\n", record[0])
		text, _ := reader.ReadString('\n')
		if strings.TrimRight(text, "\n") == record[1] {
			rightAnswers++
			fmt.Printf("Input: %sAnswer: %s\n", text, record[1])
		} else {
			wrongAnswers++
			fmt.Printf("Input: %sAnswer: %s\n", text, record[1])
		}
		total++
	}
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("\n --------- \n\nRight: %d\nWrong: %d\nTotal: %d\nTime: %v", rightAnswers, wrongAnswers, total, elapsed)
}