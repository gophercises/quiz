package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Problem structure
type Problem struct {
	q string
	a int
}

// PROBLEMBUFCOUNT sets number of problems in buffer
const PROBLEMBUFCOUNT = 100

// DEFAULTTIMELIMIT sets the default time limit for the quiz
const DEFAULTTIMELIMIT = 30 //seconds

// Problem counter
var count int

// Score counter
var score int
var faults int

// readProblems takes reads problems from file, line by line
// problems are written to problems channel
func readProblems(problems chan Problem, filename string, shuffle bool) {
	// problems buffer
	buf := make([]Problem, PROBLEMBUFCOUNT)

	// Open problems file
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

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

		// store problem in buffer
		buf[count] = Problem{q, ans}
		count++
	}

	// shuffle problems
	if shuffle {
		rand.Seed(time.Now().Unix())
		for i := range buf[:count] {
			j := rand.Intn(i + 1)
			buf[i], buf[j] = buf[j], buf[i]
		}
	}

	// send problems over channel
	for _, p := range buf[:count] {
		problems <- p
	}

}

// solveProblem consumes problems from chan
// will check timer before giving points
func solveProblem(problems chan Problem) {
	// IO reader
	scanner := bufio.NewScanner(os.Stdin)

	// start consuming problems from channel
	for p := range problems {

		// Print problem question
		fmt.Printf("%s = ", p.q)

		// Scan IO
		scanner.Scan()
		input := strings.Trim(scanner.Text(), " ") // remove whitespace
		// convert string answer to integer
		givenAns, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("'%s' is not a valid answer\n", input)
			continue
		}

		// Check answer and give points
		if givenAns == p.a {
			score++
			fmt.Println("Correct!")
		} else {
			faults++
			fmt.Println("Wrong!")
		}
	}
	return
}

// startTimer blocks for n seconds
func startTimer(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	// Parse Command line flags
	filenamePtr := flag.String("f", "problems.csv", "name of problem csv file")
	secondsPtr := flag.Int("t", DEFAULTTIMELIMIT, "number of seconds to solve problems")
	shufflePtr := flag.Bool("s", false, "shuffle questions")
	debugPtr := flag.Bool("debug", false, "show debug information")
	flag.Parse()

	// Show debug/logging information
	if !*debugPtr {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("debug:", *debugPtr)
	log.Println("filename:", *filenamePtr)
	log.Println("shuffle:", *shufflePtr)
	log.Println("timer:", *secondsPtr)

	// Problems channel (buffered)
	problems := make(chan Problem, PROBLEMBUFCOUNT)

	// read problems goroutine
	go readProblems(problems, *filenamePtr, *shufflePtr)

	// Prompt to start the game
	fmt.Printf("Press any key to start the quiz!")
	bufio.NewScanner(os.Stdin).Scan()

	// solve problems goroutine
	go solveProblem(problems)

	// start timer barrier
	startTimer(*secondsPtr)

	// Show final tally
	fmt.Printf("\nNumber of questions: %d\n", count)
	fmt.Printf("Correct answers: %d\n", score)
	fmt.Printf("Incorrect answers: %d\n", faults)

}
