package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func sanitize(s string) string {
	return strings.ToLower(strings.Trim(s, "\n\r\t"))
}

func main() {
	problemsFile := flag.String("path", "problems.csv", "This is the flag to the CSV containing the problems for the quiz")
	timeout := flag.Int("timeout", 30, "The amount of time you have for a single question in seconds")
	flag.Parse()

	file, err := os.Open(*problemsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	problems, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	correct := 0
	b := false

	for c, q := range problems {
		fmt.Printf("Question %v: %v -> ", c, q[0])

		inChan, outChan := make(chan int, 1), make(chan int, 1)
		go func() {
			inReader := bufio.NewReader(os.Stdin)
			inp, _ := inReader.ReadString('\n')
			if sanitize(inp) == sanitize(q[1]) {
				outChan <- <-inChan + 1
				return
			}
			outChan <- <-inChan
		}()
		inChan <- correct

		select {
		case res := <-outChan:
			correct = res
		case <-time.After(time.Duration(*timeout) * time.Second):
			close(outChan)
			b = true
			fmt.Println()
		}

		if b {
			break
		}

	}
	fmt.Println("You got", correct, "out of", len(problems), "correct.")
}
