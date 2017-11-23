package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Problem struct {
	q string
	a int
}

func readProblems(problems chan Problem, fd *os.File) {
	// Scan file
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		// Read comma separated line
		line := strings.Split(scanner.Text(), ",")
		// first value is the question string
		q := line[0]
		// Convert string answer to integer
		ans, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		// add problem to queue
		problems <- Problem{q, ans}
	}
	problems <- Problem{"", 0}
}

// solveProblem consumes problems from chan
// will check timer before giving points
func solveProblem(problems chan Problem, timer chan bool) (score int) {
	// Score
	score = 0
	// IO reader
	scanner := bufio.NewScanner(os.Stdin)

	// start consuming problems from channel
	for p := range problems {

		// check if all problems received (hacky)
		if p.q == "" {
			log.Println("no more problems on channel")
			break
		}

		// Print problem question
		fmt.Printf("%s = ", p.q)

		// Scan IO
		scanner.Scan()
		// convert string answer to integer
		givenAns, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal()
		}

		// Check if time-out
		select {

		case <-timer: // if time-out received, return current score
			log.Println("time out")
			return score

		default: // else, do nothing
		}

		// Check answer and give points
		if givenAns == p.a {
			score++
			fmt.Println("Correct!")
		} else {
			score--
			fmt.Println("Wrong!")
		}
	}

	return score
}

// startTimer pings timer channel after n seconds
func startTimer(seconds int, timer chan bool) {
	// sleep for n seconds
	time.Sleep(time.Duration(seconds) * time.Second)
	// write time-out to channel
	timer <- true
}

func main() {
	// Parse Command line flags
	filenamePtr := flag.String("f", "problems.csv", "name of problem csv file")
	secondsPtr := flag.Int("t", 10, "number of seconds to solve problems")
	debugPtr := flag.Bool("debug", false, "show debug information")
	flag.Parse()

	// Show debug/logging information
	if !*debugPtr {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("debug:", *debugPtr)
	log.Println("filename:", *filenamePtr)
	log.Println("timer:", *secondsPtr)

	// Open problems file
	fd, err := os.Open(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}
	// Close file before exiting
	defer fd.Close()

	// Problem queue
	problems := make(chan Problem)

	// async read problems to queue
	go readProblems(problems, fd)

	// start timer
	timer := make(chan bool)
	go startTimer(*secondsPtr, timer)

	// solve problems
	score := solveProblem(problems, timer)

	// Show final score
	fmt.Printf("Your score is %d\n", score)
}
