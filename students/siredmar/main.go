package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var (
	file  = flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	limit = flag.Int("limit", 30, "the time limit for the quiz in seconds")
)

func main() {
	flag.Parse()

	if *limit < 0 {
		fmt.Println("Error: enter positive time limit")
		os.Exit(1)
	}

	records, err := read(*file)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println("Press enter to start. Time limit is", *limit, "seconds")
	in := bufio.NewReader(os.Stdin)
	in.ReadString('\n')

	input := make(chan string)

	quizdone := make(chan bool)
	timeout := time.NewTicker(time.Duration(*limit) * time.Second)

	go getInput(input)

	var score int
	var maxscore int

	go func() {
		maxscore = len(records)
		for i, v := range records {
			fmt.Print("Problem #", i+1, ": ", v[0], " = ")
			text := <-input
			if text == v[1] {
				score++
			}
		}
		quizdone <- true
	}()

	select {
	case <-quizdone:
	case <-timeout.C:
		fmt.Println("\nThe time is up! Game over!")

	}

	fmt.Println("You scored", score, "out of", maxscore)
}

func read(filename string) ([][]string, error) {
	dat, err := ioutil.ReadFile(filename)
	r := csv.NewReader(strings.NewReader(string(dat)))
	if err == nil {
		var records [][]string
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)
		}
		return records, nil
	}
	return nil, err
}

func trim(s string) string {
	t := strings.Trim(s, "\n")
	t = strings.Trim(t, " ")
	return t
}

func getInput(input chan<- string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		result = trim(result)
		input <- result
	}
}
