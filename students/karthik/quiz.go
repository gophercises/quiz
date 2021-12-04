package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of question and answer")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open file %s \n", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("can not parse the file"))
	}
	problems := parse_csv(lines)
	var answer string
	correct := 0
	for index, pro := range problems {
		fmt.Printf("problem # %d : %s \n", index+1, pro.q)
		fmt.Scanf("%s\n", &answer)
		if answer == pro.a {
			correct++
		}
	}
	fmt.Printf("%d out of %d are correct\n", correct, len(problems))
}

func parse_csv(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for index, line := range lines {
		problems[index] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
