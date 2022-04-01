package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type Problem struct {
	Prompt string
	Answer string
}

func main() {

	timelimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	//getproblems
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problemList := createProblemList(data)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please answer the prompted questions and press Enter to begin : ")
	scanner.Scan()
	testStart := scanner.Text()
	if testStart == "" {
		fmt.Println("timer started")

	}

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	total := 0

	for _, v := range problemList {
		fmt.Println(v.Prompt)
		answerCh := make(chan string)

		go func() {
			scanner.Scan()
			ans1 := scanner.Text()
			answerCh <- ans1
		}()

		select {
		case <-timer.C:
			color.HiRed("**TIMES UP***\n")
			fmt.Printf("you answered %d correct of 12. ", total)
			return
		case ans1 := <-answerCh:
			//compare func
			if err != nil {
				fmt.Println("error converting string")
			}
			if ans1 == v.Answer {
				color.Green("correct")
				total++
			} else {
				color.Red("incorrect")
			}
		}

	}

}

//func gameTimer()

func createProblemList(data [][]string) []Problem {
	var problemList []Problem
	for i, line := range data {
		if i > 0 {
			var rec Problem
			for j, field := range line {
				if j == 0 {
					rec.Prompt = field
				} else if j == 1 {
					rec.Answer = field
				}
			}
			problemList = append(problemList, rec)
		}

	}
	return problemList
}
