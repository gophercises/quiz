package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	var csvfilename = flag.String("file", "problems.csv", "the file with the questions")
	flag.Parse()

	//  this will print out the file name that the user has given to the flag
	// by default it is problems.csv
	fmt.Println(*csvfilename)

	file, err := os.Open(*csvfilename)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)
	rec, err := r.ReadAll()
	if err != nil {
		panic("can not read the file")
	}

	ans := []string{}
	answers := []string{}
	var ans1 string

	for _, v := range rec {
		fmt.Println("solve this", v[0])
		_, err := fmt.Scanln(&ans1)
		if err != nil {
			panic(err)
		}
		ans = append(ans, ans1)

		for i := 0; i < len(rec); i++ {

			answers = append(answers, string(rec[i][1]))
		}
	}
	correctans := 0

	for i := 0; i < len(ans); i++ {
		if ans[i] == answers[i] {
			correctans += 1
		}
	}
	fmt.Println("Number of correct answers are :", correctans)

}
