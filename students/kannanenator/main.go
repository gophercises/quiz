package main

import "os"
import "fmt"
import "log"
import "flag"
import "time"
import "encoding/csv"
import "bufio"
import "strings"

func main() {
	
	filenamePtr := flag.String("filename", "problems.csv", "file containing the set of problems")
	limitPtr := flag.Int("limit", 30, "quiz time limit")

	flag.Parse()

	file, err := os.Open(*filenamePtr)
	handleError(err)
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	handleError(err)
	
	numQs := len(rows)
	numCorrect := 0

	timer := time.NewTimer(time.Second * time.Duration(*limitPtr))
	go func() {
		<- timer.C
		// when the timer ends, we kill the quiz
		fmt.Println("\nTime is up")
		os.Exit(0)
	}()
	
	consoleReader := bufio.NewReader(os.Stdin)
	for idx, element := range rows {
		q, a := element[0], element[1]
		fmt.Print("Problem #", idx+1 ,": ", q, " = ")
		input, _ := consoleReader.ReadString('\n')
		
		// compare w/o whitespace
		if strings.TrimSpace(input) == strings.TrimSpace(a) {
			numCorrect++
		}
	}

	fmt.Println("You got", numCorrect, "out of", numQs, "correct")
}

func handleError(err error){
	if err != nil {
		log.Fatal(err)
	}
}
