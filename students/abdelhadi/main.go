package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var FullScore int
var UserScore int
var Coninue bool

func init() {
	FullScore = 12
	UserScore = 0
	Coninue = true
}

func PrintUserScore() {
	fmt.Printf("\nYou Score %v out of %v:\n", UserScore, FullScore)
	os.Exit(0)
}

func main() {

	csvFileName := flag.String("f", "wrongfile.csv", "csv file name") //the true file name file.csv
	AnswerTime := flag.Int("t", 2, "Answer Time in secounds")
	flag.Parse()
	csvFile, err := os.Open(*csvFileName)
	defer csvFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(csvFile)
scanFileLoop:
	for scanner.Scan() {
		textLine := scanner.Text()
		splits := strings.Split(textLine, ",")
		abordCh := make(chan bool)
		var userAnswer string
		fmt.Printf("%v:", splits[0])

		go contiuoApp(abordCh, *AnswerTime)

		if _, err := fmt.Scan(&userAnswer); err != nil {
			log.Fatal(err)
		}

		answerResult := quiz(splits[1], userAnswer)
		if answerResult == true {
			UserScore++
			//contiuo the app
			abordCh <- false
		} else if answerResult == false {
			//call func is enouph
			PrintUserScore()
			//Abord The App
			abordCh <- true
			break scanFileLoop
		}
	}
}

func contiuoApp(ch chan bool, AnswerTime int) {
	ticker := time.NewTicker(time.Second * time.Duration(AnswerTime))
	go func() {
		select {
		case <-ch:
		case <-ticker.C:
			PrintUserScore() //and exit()
		}
	}()
}

func quiz(answer string, userAnswer string) bool {
	if userAnswer == answer {
		return true
	}
	return false
}
