package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const problemsFilename = "problems.csv"

func main() {

	//var myfile *file It will effect all file that read myfile
	// var myfile file It will return only copy of a file will not effect on original file

	f, err := os.Open(problemsFilename)
	if err != nil {
		fmt.Printf("failed to open file : %v", err)
		return
	}

	//defer is used to clean after I finish the recurse
	//After main package is finished It will return back
	defer f.Close()
	// read csv file (problem.csv).
	read := csv.NewReader(f)
	records, err := read.ReadAll()

	if err != nil {
		fmt.Printf("failed to read csv file : %v\n", err)
		return
	}

	// range iterate will output 2 values index , value
	var correctAnswers int
	for i, record := range records {
		//fmt.Println(i)
		//fmt.Println(record)
		question, correctAnswer := record[0], record[1]
		//display one question at a time.
		fmt.Printf("%d. %s?\n", i+1, question)
		//get answer from user, then proceed to next question
		//immediately
		var answer string

		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Printf("failed to scan :%v\n", err)
			return
		}
		fmt.Printf("Your answer: %s\n", answer)

		if answer == correctAnswer {
			correctAnswers++
		}
	}

	//output number of questions(total + correct)
	fmt.Printf("Result : %d%d", correctAnswers, len(records))

}

//fmt.Println(r)
