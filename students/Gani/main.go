package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	csvPath = flag.String("csv", "problem.csv", "This is string")
	limit   = flag.Int("limit", 30, "This works as time for the ticker")
)

func main() {

	flag.Parse()

	file, err := os.Open(*csvPath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	//open the csv file

	csvReader := csv.NewReader(file)

	csvData, err := csvReader.ReadAll()

	if err != nil {
		panic(err)
	}

	qaPair := make(map[string]string, len(csvData))

	for _, data := range csvData {
		qaPair[data[0]] = data[1]
	}

	ticker := time.NewTicker(time.Second * time.Duration(*limit))

	done := make(chan bool)

	scanner := bufio.NewScanner(os.Stdin)

	var userAnswer string

	qna, correct := 0, 0

	go func() {

		for question, answer := range qaPair {
			qna++
			fmt.Printf("Problem #%d: %s = ", qna, question)
			scanner.Scan()
			userAnswer = scanner.Text()

			userAnswer = strings.TrimSpace(userAnswer)
			userAnswer = strings.ToLower(userAnswer)

			if userAnswer == answer{
				correct++;
			}


		}
		done <- true

	}()


	select{
	case <-done:
	case <- ticker.C:
		fmt.Println("Times up!")
	}

	fmt.Printf("You scored %d out of %d ", correct, len(csvData))

}
