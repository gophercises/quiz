package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	fileName := flag.String("csv", "problems.csv", "give the name to csv file containing problems")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(file)
	questionAndAnswers, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	quiz := addToDataStructure(questionAndAnswers)
	// fmt.Println(quiz)

	t := time.NewTimer(time.Second * 30)

	totalMarks := 0
	for i, q := range quiz {
		fmt.Printf("question no:%d question:%s options:%v\n", i+1, q.problem, q.choice)
		ansC := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			ansC <- ans
		}()

		select {
		case <-t.C:
			fmt.Printf("total marks scored by the player : %d", totalMarks)
			return
		case answer := <-ansC:
			if answer == q.answer {
				totalMarks++
			}
		}
	}

	fmt.Printf("total marks scored = %d\n", totalMarks)

}

type Quiz struct {
	problem string
	choice  []string
	answer  string
}

func addToDataStructure(questionAndAnswers [][]string) []Quiz {

	quiz := make([]Quiz, len(questionAndAnswers))

	for i, val := range questionAndAnswers {

		quiz[i] = Quiz{
			problem: val[0],
			choice:  val[1 : len(val)-1],
			answer:  val[len(val)-1],
		}
	}

	return quiz
}
