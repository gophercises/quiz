package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	csvPath = flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit   = flag.Int("limit", 20, "the time limit for the quiz in seconds'")
)

func main() {
	flag.Parse()
	file, err := openFile(*csvPath)

	if err != nil {
		panic(err.Error())
	}

	statements := readCSV(file)
	scanner := bufio.NewScanner(os.Stdin)

	done := make(chan bool)
	// create a ticker for the time limit and a channel to signal the user finished the quiz
	ticker := time.NewTicker(time.Second * time.Duration(*limit))
	var score int
	go func() {
		for i, v := range statements {
			fmt.Printf("#%d: %s = ", i+1, v.q)
			_ = scanner.Scan()
			answer := scanner.Text()
			if answer == v.a {
				score++
			}
		}
		done <- true
	}()

	select {
	case <-done:
	case <-ticker.C:
		fmt.Println("Time is up buddy!")
	}
	fmt.Printf("Scored %d out to 12", score)
}

// open the file, returning the interface
func openFile(fileName string) (io.Reader, error) {
	return os.Open(fileName)
}

func readCSV(r io.Reader) []statement {
	recs, err := csv.NewReader(r).ReadAll()

	if err != nil {
		log.Println("cannot read reader, do'h")
		panic(err.Error())
	}

	total := len(recs)

	log.Printf("total: %d", total)

	var ss []statement
	for _, v := range recs {
		var s statement
		s.q = v[0]
		s.a = v[1]
		ss = append(ss, s)
	}

	return ss
}

type statement struct {
	q string
	a string
}
