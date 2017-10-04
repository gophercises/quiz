package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	flagFilePath string
	flagRandom   bool
	flagTime     int
	wg           sync.WaitGroup
)

func init() {
	flag.StringVar(&flagFilePath, "file", "questions.csv", "path/to/csv_file")
	flag.BoolVar(&flagRandom, "random", true, "randomize order of questions")
	flag.IntVar(&flagTime, "time", 10, "test duration")
	flag.Parse()
}

func main() {
	// this program will progress as follows
	// read a csv filepath and a time limit from flags
	// prompt for a key press
	// on key press, start the quiz as follows
	//
	// while time has not elapsed:
	// print a random question to the screen
	// prompt the user for an answer
	// store the answer in a container
	// normalize answers so they compare correctly
	// output total questions answered correctly and how many questions there
	// were.

	csvPath, err := filepath.Abs(flagFilePath)
	if err != nil {
		log.Fatalln("Unable to parse path" + csvPath)
	}
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var totalQuestions = len(csvData)
	questions := make(map[int]string, totalQuestions)
	answers := make(map[int]string, totalQuestions)
	responses := make(map[int]string, totalQuestions)

	for i, data := range csvData {
		questions[i] = data[0]
		answers[i] = data[1]
	}

	respondTo := make(chan string)

	// block until user presses enter
	fmt.Println("Press [Enter] to start test.")
	bufio.NewScanner(os.Stdout).Scan()
	if flagRandom {
		// seed the random number generator with the current time
		rand.Seed(time.Now().UTC().UnixNano())
	}
	// randPool should contain random indexes into the questions map
	randPool := rand.Perm(totalQuestions)

	wg.Add(1)
	timeUp := time.After(time.Second * time.Duration(flagTime))
	go func() {
	label:
		for i := 0; i < totalQuestions; i++ {
			index := randPool[i]
			go askQuestion(os.Stdout, os.Stdin, questions[index], respondTo)
			select {
			case <-timeUp:
				fmt.Fprintln(os.Stderr, "\nTime up!")
				break label
			case ans, ok := <-respondTo:
				if ok {
					responses[index] = ans
				} else {
					break label
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()

	correct := 0
	for i := 0; i < totalQuestions; i++ {
		if checkAnswer(answers[i], responses[i]) {
			correct++
		}
	}
	summary(correct, totalQuestions)
}

func askQuestion(w io.Writer, r io.Reader, question string, replyTo chan string) {
	reader := bufio.NewReader(r)
	fmt.Fprintln(w, "Question: "+question)
	fmt.Fprint(w, "Answer: ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		close(replyTo)
		if err == io.EOF {
			return
		}
		log.Fatalln(err)
	}
	replyTo <- answer
}

func checkAnswer(ans string, expected string) bool {
	if strings.EqualFold(strings.TrimSpace(ans), strings.TrimSpace(expected)) {
		return true
	}
	return false
}

func summary(correct, totalQuestions int) {
	fmt.Fprintf(os.Stdout, "You answered %d questions correctly (%d / %d)\n", correct,
		correct, totalQuestions)
}
