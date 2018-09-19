package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const timeToAnswer = 5 * time.Second

type quiz struct {
	challenge, response string
}

func (q *quiz) ask(timeout time.Duration, lines <-chan string, roundOver chan<- struct{}) bool {
	color.Set(color.FgGreen)
	fmt.Print("> ")
	color.Unset()
	fmt.Println(q.challenge)

	select {
	case line := <-lines:
		return clean(line) == clean(q.response)
	case <-time.After(timeout):
		color.Set(color.FgRed)
		fmt.Println("Out of time")
		color.Unset()

		roundOver <- struct{}{}

		return false
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`USAGE: go run main.go <CSV_FILE>`)
		fmt.Println(len(os.Args))
		return
	}

	qs, err := readCSV(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to read file %s: %v\n", os.Args[1], err)
		return
	}

	lines := make(chan string)
	roundOver := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), timeToAnswer*time.Duration(len(qs)))

	go listenForUserInput(ctx, bufio.NewReader(os.Stdin), lines, roundOver)

	good := 0
	for _, q := range qs {
		if q.ask(timeToAnswer, lines, roundOver) {
			good++
		}
	}

	cancel()

	fmt.Printf("Answered %d/%d questions.\n", good, len(qs))
}

func clean(a string) string {
	return strings.TrimSpace(strings.ToLower(a))
}

func listenForUserInput(ctx context.Context, r io.RuneReader, lines chan<- string, roundOver <-chan struct{}) {
	inputRunes := []rune{}
	newline := false

	for {
		select {
		case <-ctx.Done():
			close(lines)
			return
		case <-roundOver:
			inputRunes = inputRunes[:0]
			newline = false
		default:
			if newline {
				lines <- string(inputRunes)
				inputRunes = inputRunes[:0]
				newline = false
			}

			run, _, err := r.ReadRune()
			if err != nil {
				log.Fatalln("Couldn't read rune:", err)
			}

			if run == '\n' {
				newline = true
			} else {
				inputRunes = append(inputRunes, run)
			}
		}
	}
}

func readCSV(filename string) ([]quiz, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer f.Close() // nolint

	r := csv.NewReader(f)
	out := []quiz{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error while reading CSV record: %v", err)
		}
		if len(record) != 2 {
			return nil, fmt.Errorf("unexpected number of fields for record: %v", record)
		}

		out = append(out, quiz{record[0], record[1]})
	}

	return out, nil
}
