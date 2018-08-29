package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"
)

func startTime(timer time.Timer, duration time.Duration) {
	<-timer.C
	fmt.Println(duration, "doldu")
}

func main() {
	filename := flag.String("file", "problems.csv", "Questions file")
	timelimit := flag.Int("time", 30, "Time limit of quiz in seconds")
	flag.Parse()

	filePath, pathErr := filepath.Abs(*filename)
	if pathErr != nil {
		log.Fatal("file not found. check if it exists again.")
	}

	remainTime := time.Duration(*timelimit) * time.Second
	fmt.Println("Press return to start time (", remainTime, ")")
	fmt.Scanln() //and time starts

	timer := time.NewTimer(remainTime)
	go startTime(*timer, remainTime)

	csvF, csvErr := ioutil.ReadFile(filePath)
	if csvErr != nil {
		log.Fatal("error reading file: " + csvErr.Error())
	}

	probs, probErr := csv.NewReader(strings.NewReader(string(csvF))).ReadAll()
	if probErr != nil {
		log.Fatal("error on CSV format: " + probErr.Error())
	}

	var answer string
	var trues, total int

	for i, soru := range probs {
		fmt.Printf("%d. soru: %s = ", i+1, soru[0])
		fmt.Scan(&answer)
		if answer == soru[1] {
			trues++
		}
		total++
	}

	fmt.Printf("True answers: %d\n", trues)
	fmt.Printf("Total questions: %d\n", total)
}
