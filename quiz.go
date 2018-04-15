package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var timelimit = flag.Int("tl", 30, "Timelimit for quiz")
var shuffle = flag.Int("shuffle", 0, "Pass in 1 to shuffle")

func main() {
	flag.Parse()

	problems := readData()
	if *shuffle == 1 {
		problems = shuffleProblems(problems)
	}

	promptStart()
	runQuiz(problems)
}

func readData() [][]string {
	f, _ := os.Open("./problems.csv")
	r := csv.NewReader(io.Reader(f))

	problems, _ := r.ReadAll()
	return problems
}

func promptStart() {
	fmt.Println("Press ENTER to start...")
	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')
}

func shuffleProblems(data [][]string) [][]string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuf := make([][]string, len(data))
	perm := r.Perm(len(data))
	for i, randIndex := range perm {
		shuf[i] = data[randIndex]
	}

	return shuf
}

func runQuiz(problems [][]string) {
	correct := 0
	total := len(problems)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
		<-timer.C
		fmt.Printf("\nTime's up!\n")
		wg.Done()
	}()

	go func() {
		for _, problem := range problems {
			question := problem[0]
			answer := problem[1]
			fmt.Printf("%s: ", question)

			r := bufio.NewReader(os.Stdin)
			guess, _ := r.ReadString('\n')
			guess = strings.TrimSpace(guess)
			if strings.Compare(strings.ToLower(answer), strings.ToLower(guess)) == 0 {
				correct++
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("You answered %d of %d correct", correct, total)
}
