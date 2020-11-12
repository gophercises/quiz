package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("./problems.csv")
	check(err)
	defer file.Close()

	var timeLimit int
	flag.IntVar(&timeLimit, "limit", 0, "time limit")
	flag.Parse()

	score := 0
	csvReader := csv.NewReader(file)
	buf := bufio.NewReader(os.Stdin)
	fmt.Printf("Time limit is %d seconds, press <ENTER> to start!", timeLimit)
	fmt.Scanln()

	done := make(chan bool)

	go func() {
		for {
			quiz, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			// print quiz
			fmt.Printf(quiz[0] + ": ")

			// read answer
			answer, err := buf.ReadString('\n')
			if err != nil {
				answer = ""
			}
			if strings.TrimRight(answer, "\r\n") == quiz[1] {
				score++
			}
		}
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("You did it!")
	case <-time.After(time.Duration(timeLimit) * time.Second):
		fmt.Println("Ok, time is up!")
	}
	fmt.Printf("Your score is %d\n", score)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
