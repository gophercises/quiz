package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type entry struct {
	question string
	answer   string
}

var (
	srcFile  = flag.String("source", "problems.csv", "A CSV file containing problems and answers, separated by commas")
	timeout  = flag.Int("timeout", 30, "The time in seconds that is alloted to solve as many questions as possible")
	shuffle  = flag.Bool("shuffle", false, "If the questions in the source should be shuffled")
	nbrEntry int
	answered int
	correct  int
)

func readSrcFile(csvFile *os.File) []entry {
	r := csv.NewReader(csvFile)

	entries := make([]entry, 0)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Could not read record")
		}
		entries = append(entries, entry{strings.TrimSpace(rec[0]), strings.TrimSpace(rec[1])})
	}
	return entries
}

func shuffleSlice(entries []entry) {
	rand.Seed(time.Now().UTC().UnixNano())
	rand.Shuffle(len(entries), func(i, j int) {
		entries[i], entries[j] = entries[j], entries[i]
	})
}

func printStats() {
	fmt.Printf("Total: %d, Answered: %d, Correct: %d, Wrong: %d\n", nbrEntry, answered, correct, answered-correct)
}

func printInstructions() {
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	fmt.Printf("# Welcome to the quiz game, press a key to start #\n")
	fmt.Printf("#                                                #\n")
	fmt.Printf("#     You have %d seconds to finish the game     #\n", *timeout)
	fmt.Printf("%s\n\n", strings.Repeat("=", 50))
}

func main() {
	flag.Parse()

	f, err := os.Open(*srcFile)
	if err != nil {
		fmt.Println("Could not open file")
		return
	}
	defer f.Close()

	entries := readSrcFile(f)
	nbrEntry = len(entries)

	if *shuffle {
		shuffleSlice(entries)
	}

	printInstructions()
	fmt.Scanln()

	timeLimit := time.NewTicker(time.Second * time.Duration(*timeout))

	for _, entry := range entries {
		c1 := make(chan string, 1)

		go func() {
			fmt.Printf("%s = ?: ", entry.question)
			var s string
			fmt.Scanln(&s)
			c1 <- strings.TrimSpace(s)
		}()
		select {
		case <-timeLimit.C:
			fmt.Printf("\nTime has run out\n")
			printStats()
			return
		case guess := <-c1:
			if guess == entry.answer {
				correct++
			}
			answered++
		}
	}

	printStats()
}
