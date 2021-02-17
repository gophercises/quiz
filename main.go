package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	shuffled := flag.Bool("shuffled", false, "Should the questions be shuffled?")
	limit := flag.Int("limit", 30, "Time limit in seconds.")
	flag.Parse()
	lines, err := readlines("problems.csv")
	questions := len(lines)
	if err != nil {
		return
	}
	if *shuffled == true {
		lines = shuffle(lines)
	}
	quiz(lines, limit, questions)
}

func readlines(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func shuffle(stringarray []string) []string {
	shuffled := make([]string, len(stringarray))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(stringarray))
	for i, v := range perm {
		shuffled[v] = stringarray[i]
	}
	return shuffled
}

func quiz(lines []string, limit *int, questions int) {
	var correct int = 0
	var guess string
	fmt.Print("Press enter to begin the quiz:")
	fmt.Scanln()
	ticker := time.NewTicker(time.Duration(*limit) * time.Second)
	timesup := make(chan bool)
	go func() {
		for i := 0; i < questions; i++ {
			lines[i] = strings.ReplaceAll(lines[i], ", ", " ")
			question := lines[i][:strings.IndexByte(lines[i], ',')]
			answer := lines[i][strings.IndexByte(lines[i], ',')+1:]
			fmt.Print(question, " = ")
			fmt.Scanln(&guess)
			guess = strings.ToLower(strings.TrimSpace(guess))
			if guess == answer {
				correct++
			}
		}
		timesup <- true
	}()
	select {
	case <-timesup:
	case <-ticker.C:
	}
	fmt.Printf("\nYou got %d out of %d questions correct.\n", correct, questions)
}
