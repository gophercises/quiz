package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	file, _ := os.Open("./problems.csv")
	r := csv.NewReader(io.Reader(file))

	problems, _ := r.ReadAll()

	fmt.Println("Press ENTER to start...")
	startReader := bufio.NewReader(os.Stdin)
	startReader.ReadString('\n')
	runQuiz(problems)
}

func runQuiz(problems [][]string) {
	correct := 0
	total := len(problems)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		timer := time.NewTimer(30 * time.Second)
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
			if strings.Compare(answer, guess) == 0 {
				correct++
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("You answered %d of %d correct", correct, total)
}
