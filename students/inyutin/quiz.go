package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Line struct {
	Question string
	Answer   string
}

var (
	csvName = flag.String("csv", "problems.csv", "path to csv file with quiz(question,answer)")
	limit   = flag.Int("limit", 30, "the time limit for the quiz in seconds")
)

func main() {
	flag.Parse()

	csvFile, err := os.Open(*csvName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	var lines []Line
	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		lines = append(lines, Line{
			Question: line[0],
			Answer:   line[1],
		})
	}

	reader := bufio.NewReader(os.Stdin)
	count := 0

	T := time.Duration(*limit)
	timer := time.NewTimer(T * time.Second)
	go func() {
		<-timer.C
		fmt.Println()
		fmt.Println("You scored " + strconv.Itoa(count) + " out of " + strconv.Itoa(len(lines)))
		os.Exit(0)
	}()

	for idx, line := range lines {
		fmt.Print("Question №" + strconv.Itoa(idx+1) + ": " + line.Question + " = ")
		ans, _ := reader.ReadString('\n')
		if ans == line.Answer+"\n" {
			count++
		}
	}
	stop := timer.Stop()
	if stop {
		fmt.Println("You scored " + strconv.Itoa(count) + " out of " + strconv.Itoa(len(lines)))
	}
}
