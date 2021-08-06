package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//file flag
var fileName = flag.String("file", "problems.csv", "the csv file containing questions")
//time flag
var timeLimit = flag.Int("limit" ,30 , "timer for each problem")

//getting response for each question
func getResponse(reader *bufio.Reader , input chan<- string){ //this type of channel is passed when we need to pass something to it
	resp ,err  := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	input <- strings.Replace(resp , "\n" , "" , -1)
}

//asks the questions and gets the response withing a time limit
func askQuestions(reader *csv.Reader) (responses []bool){
	
	problems , _:= reader.ReadAll()

	responses = make([]bool , len(problems))

	//so as to get the responses
	stdReader := bufio.NewReader(os.Stdin)

	//started the timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//make a channel for string
	input := make(chan string)

	problemLoop: //label
	//for each problem
		for index , problem := range problems{

			question , answer := problem[0] , strings.TrimSpace(problem[1]) //to keep the trailing and the beginning spaces from the answer

			fmt.Printf("Problem #%d : %s = " , index + 1 , question)

			//make a separate go routine
			//reason being the ReadString will freeze to the point so even if the timer had run out , it would still keep going until
			//we give the required input

			go getResponse(stdReader , input)

			select{
			case <-timer.C:
				fmt.Println()
				break problemLoop
			case resp := <-input:
				if resp == answer{
					responses[index] = true
				}
				break
			}
		}
	return responses
}

//calculates the result
func calculateResult(responses []bool) int{
	var count = 0
	for _ ,response := range responses{
		if response{
			count++
		}
	}
	return count
}

//gets the file pointer
func getFile(filename string) (*os.File , error){
	file , err := os.Open(filename)
	if err != nil{
		return nil , errors.New("Not able to open the file")
	}
	return file, nil
}

//returns a csv reader to the file specified in the global context
func getCsvReader() ( *csv.Reader , error){
	file , err := getFile(*fileName)
	if err != nil{
		return nil , err
	}
	return csv.NewReader(file) , nil
}

func main() {

	//start the program by first parsing through the flags
	//contains the flags for the filename and the time limit - contains default valeus too
	//filename := "problems.csv" , limit := 30 sec
	flag.Parse()

	// fmt.Println(*fileName)

	//get the csv reader for the file
	reader,err := getCsvReader()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("There will be a time limit of %d sec for each question .\n" , *timeLimit)
	fmt.Println("Press Enter to continue.")
	
	bufio.NewReader(os.Stdin).ReadBytes('\n') //prompt to start the timer

	fmt.Println("Asking the questions now ...")
	//asks the questions and gets the responses
	responses := askQuestions(reader)
	//calculates the result
	result := calculateResult(responses)

	fmt.Printf("Scored %d out of %d\n" , result , len(responses))
}