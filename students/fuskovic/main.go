package main

import (
	"encoding/csv"
	"os"
	"io"
	"fmt"
	"bufio"
	"strconv"
	"flag"
	"time"
	"strings"
)

type question struct {
	question	string
	answer		int
}

var (
	quiz []question
	t = flag.Int("timer",30,"Allotment of time in seconds to complete the quiz.")
	p = flag.String("file","problems.csv","Filepath of csv containing questions and answers.")
	mix = flag.Bool("shuffle",false,"Determine if questions should be shuffled.")
	correct = 0
)

func startTimer(startTime time.Time , endTime *int, ch chan<- string){
	for {
		elapsed := time.Since(startTime).Seconds()
		if int(elapsed) == *endTime{
			ch <- "stop"
			break
		}
	}
}

func watchTimer(ch <-chan string){
	for {
		if <-ch == "stop"{
			fmt.Printf("\ntimes up\nresults : %d/%d correct\n",correct, len(quiz))
			os.Exit(1)
		}
	}
}

func shuffle(answerSheet []question) []question{
	qAndA := make(map[string]int)
	for _, question := range answerSheet {
		qAndA[question.question] = question.answer
	}
	var shuffledQuiz []question
	for q,a := range qAndA {
		shuffledQuiz = append(shuffledQuiz,question{
			question : q,
			answer : a,
		})
	}
	quiz = shuffledQuiz
	return quiz
}

func generateQuiz(r *csv.Reader, f *os.File){
	for {
		line, err := r.Read()
		if err == io.EOF{
			f.Close()
			break
		}else if err != nil {
			f.Close()
			fmt.Fprintf(os.Stderr, "err found : %v\n", err)
			os.Exit(1)
		}
		a,err := strconv.Atoi(line[1])
		if err != nil{
			fmt.Println(err)
			f.Close()
			os.Exit(1)
		}
		q := &question{
			question : line[0],
			answer : a,
		}
		quiz = append(quiz,*q)
	}
}

func startQuiz(){
	if *mix {
		quiz = shuffle(quiz)
	}
	for _, question := range quiz {
		fmt.Print(question.question+":")
		s := bufio.NewScanner(os.Stdin)
			for s.Scan(){
				input,_ := strconv.Atoi(strings.ToLower(strings.TrimSpace(s.Text())))
				if input == question.answer {
					correct++
				}
				break
			}
		}
	fmt.Printf("results : %d/%d correct\n",correct, len(quiz))
}

func main(){
	flag.Parse()
	fmt.Printf("Press any key to start timer of %d seconds\n",*t)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan(){
		input := s.Text()
		if len(input) >= 0{
			break
		}
	}

	f,err := os.Open(*p)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	generateQuiz(r,f)

	ch := make(chan string)
	start := time.Now()
	go startTimer(start,t,ch)
	go watchTimer(ch)
	startQuiz()
}