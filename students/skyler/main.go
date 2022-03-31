package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Problem struct {
	Prompt string
	Answer string
}

func main() {
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

	total := 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please answer the prompted questions and then press Enter: ")
	for _, v := range problemList {
		fmt.Println(v.Prompt)

		//compare func
		scanner.Scan()
		ans1 := scanner.Text()
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

	fmt.Printf("you answered %d correct of 12. ", total)

}

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
