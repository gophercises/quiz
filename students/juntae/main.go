package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

// Problem structure for each question
type Problem struct {
	question string
	answer   string
}

func main() {
	file, limit := readArguments()

	f, err := os.Open(file)
	//Error occurred
	if err != nil {
		fmt.Printf("Failed to open the file %s\n", file)
		os.Exit(1)
	}

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("failed to parse the provided CSV file")
		os.Exit(1)
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(limit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("number of correct : %d\n", correct)
			return
		case ans := <-answerCh:
			if ans == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("number of correct : %d\n", correct)

}

func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i := 0; i < len(ret); i++ {
		ret[i] = Problem{lines[i][0], lines[i][1]}
	}

	return ret
}
func readArguments() (string, int) {
	csvPtr := flag.String("csv", "./problems.csv", "CSV File that conatins quiz questions")
	limitPtr := flag.Int("limit", 5, "Time Limit for questions")
	flag.Parse()
	return *csvPtr, *limitPtr
}
