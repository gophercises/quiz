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

func solveProblem(problems chan Problem, timer chan bool) (score int) {
	// Score
	score = 0
	// IO reader
	scanner := bufio.NewScanner(os.Stdin)

	for p := range problems {
		// all problems received
		if p.q == "" {
			log.Println("done receiving problems")
			break
		}
		// Print question
		fmt.Printf("%s = ", p.q)
		// Scan IO
		scanner.Scan()
		// convert string answer to integer
		givenAns, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal()
		}

		// Check if time out
		select {
		case <-timer:
			log.Println("time out")
			return score
		default:
			break
		}

		// Check answer
		if givenAns == p.a {
			score++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Wrong!")
		}
	}

	return score
}

func startTimer(seconds int, timer chan bool) {
	time.Sleep(time.Duration(seconds) * time.Second)

	timer <- true
}

func main() {
	log.SetOutput(ioutil.Discard)
	// Command line flags
	filenamePtr := flag.String("f", "problems.csv", "name of problem csv file")
	secondsPtr := flag.Int("d", 10, "number of seconds to solve problems")
	flag.Parse()

	log.Println("filename:", *filenamePtr)
	log.Println("delay:", *secondsPtr)

	// Open file
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

	fmt.Printf("Your score is %d\n", score)
}
