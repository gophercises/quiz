package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

/*
 usage: go run main.go -limit 5 // here five time timeout in sec
 1.  Read and parse CSV
 2. using Channel to send timeout and break the loop
*/


type Data struct {
	question string `json:"question"`
	answer string  `json:"answer"`
}


func ReadCSV(csvFile string) []Data {
	var dataset []Data
	//load csv
	f, _ := os.Open(csvFile)

	//create a reader
	r := csv.NewReader(f)

	for{
		record, err := r.Read()

		if err == io.EOF {
			break;
		}
		if err != nil{
			panic("Error occurred while reading")
		}
		dataset = append(dataset, Data{record[0], record[1]})
	}
	return dataset
}


func main()  {
 dataset := ReadCSV("../../problems.csv")
 var score int = 0
 // flag used to create flag kind of stuff it is basically pointer *int
 timeLimit := flag.Int("limit", 30, "Time limit for Quiz in seconds") // this can directly given as input using -limit from arguments
 flag.Parse() // after all flag defined use this to bind

 //timelimit will int* will point to its value
 timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) //predefined function that send time expired via channel
 //<-timer.C //to block until timer send a message in channel

 for i, data := range dataset{
	 fmt.Print("Problem #",i+1,": ",data.question," = ")
	 answerChannel := make(chan string)
	 go func() {
		 var userAnswer string
		 _, _ = fmt.Scanln(&userAnswer)
		 answerChannel <- userAnswer
	 }() // () is because we are calling anonymous function

	 select { //processing channel
	 case <-timer.C:
	 	fmt.Printf("\nOpps, TimeOut ! \n You scored %d out of %d. \n",score, len(dataset))
		 return
	 case answer:= <-answerChannel:
		 if answer == data.answer{
		 score  +=1
	 }
	 }
	 }
	fmt.Println("\nYour score is :", score,"/",len(dataset))

 }


