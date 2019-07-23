package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type quiz struct {
	total     int
	correct   int
	incorrect int
	items     []*item
}
type item struct {
	question string
	answer   string
	got      string
	correct  bool
}

func main() {
	file := flag.String("file", "problems.csv", "file to parse the quiz")
	flag.Parse()
	qz, err := load(*file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	start(qz)
	score(qz)

}

// load loads the values of the file and return a new quiz
func load(file string) (*quiz, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	all, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	qz := new(quiz)
	qz.total = len(all)

	for _, items := range all {
		i := item{question: items[0], answer: items[1]}
		qz.items = append(qz.items, &i)
	}
	return qz, nil
}

// score prints the quiz score
func score(qz *quiz) {
	fmt.Printf("total: %v correct: %v incorrect: %v\n", qz.total, qz.correct, qz.incorrect)
	fmt.Println("Incorrect questions")
	for _, q := range qz.items {
		if !q.correct {
			fmt.Printf("%v answer: %v got: %v\n", q.question, q.answer, q.got)
		}
	}
}

// start starts the quiz
func start(qz *quiz) {
	fmt.Printf("interactive mode (%v questions in quiz)\n", qz.total)
	for i, item := range qz.items {
		var op = ""
		fmt.Printf("%v) %v?: ", i+1, item.question)
		fmt.Scanln(&op)
		item.got = op
		if op == item.answer {
			qz.correct++
			item.correct = true
			continue
		}
		qz.incorrect++
	}
}
