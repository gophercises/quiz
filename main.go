package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//Channels, go routines, time package

type QuizQA struct {
	question string
	answer int
}

type Score struct {
	score int
}

func main () {
csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer")
timerDuration := flag.Int("timeDuration", 30, "The duration in seconds you have to answer the questions in this quiz")
flag.Parse()

	// go quizTimer(*timerDuration)



	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
	} 
	
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

//This stops the quiz function fromrunning as it is a blocker waiting to receive
// timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)
// 	<-timer.C
//end of blocker

	//This is setup below problems to ensure no lag for the user.
	timer := time.NewTimer(time.Duration(*timerDuration) * time.Second)
	
	correct := 0

	//user input section
	for i, p := range problems{
		//Select statement to stop the for loo
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		default:
			fmt.Printf("Problem #%d:  %s = \n", i+1, p.q)
		//Allowing the user to input, scanF isn't always appropriate.scanf good for numbers not for strings. It clears all the whitespace
		var answer string
		//reference to answer
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		} 
		}
		//Start timer here with an enter input
		
		//Could put in an else statement here to do other logic or say that was incorrect etc. Give a running total update
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}


//Append function ressizes the slice, so if the amount of data is known don't use it.

	func parseLines(lines [][]string) []problem {
		ret := make([]problem, len(lines))
		for i, line := range lines {
			ret[i] = problem{
				q: line[0],
				//line below is to make sure inputted data from the CSV is correct with no whitespace. This is an extra if it was 5+5,    10 instead of 5+5,10
				a: strings.TrimSpace(line[1]),
			}
		}
		return ret
	}
	//new
	type problem struct {
		q string
		a string
	}

	func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
		}

	//Step 1 Create a 30 second timer that runs concurrently		
	//Step 2 use this timer to stop the quiz
		//I am going to use a timer function that runs concurrently with the quiz question.. Then using a channel sends a Stop method to the quiz function which then moves straight to the display.

	//My own timer	
	// func quizTimer(in int){
	// 	time.Sleep(time.Duration(in) * time.Second)
	// 	fmt.Println("Time is up")
	// }

















//     var quizQsAndAs []QuizQA
// 	for _, record := range records {
// 		answerInt, _ := strconv.Atoi(record[1])
// 		data := QuizQA{
// 			question: record[0],
// 			answer: answerInt,
// 		}
// 		quizQsAndAs = append(quizQsAndAs, data)
// 	}
// 	quizQuestions(quizQsAndAs)
// }

// //This is the function handling the error exit


// func quizQuestions(d []QuizQA) (int){
// 	score := 0
// 	reader := bufio.NewReader(os.Stdin)
// 	count := 0
// 	for count < len(d){
// 		getInput(d[count].question, d[count].answer, score, reader)
// 		count ++
// 	}
// 	fmt.Println(score)
// 	return score
// }

// //This step is to make the questions an input that the user can answer
// func getInput(prompt string, answer int, score int, r *bufio.Reader){
// 	fmt.Print(prompt, "\n")
// 	input, err := r.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	answerGiven, _ := strconv.Atoi(input)
// 	answerCheck(answerGiven, score)
// 	if(answerGiven == answer){
// 		fmt.Print("correct answer")
// 	}
// }


// func answerCheck (ans int, score int){
// 	if(ans == 11) {
// 			fmt.Print("correct answer 1")
// 	}
// }

//ATTEMPT TO USE  BUFIO FOR USER INPUT "
	// for i, p := range problems{
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Printf("Problem #%d:  %s = \n", i+1, p.q)
	// 	answer, _ := reader.ReadString('\n')
	// 	fmt.Print(answer)
	// 	if answer == p.a {
	// 		fmt.Println("correct")
	// 		correct++
	// 	} 
	// 	//Could put in an else statement here to do other logic or say that was incorrect etc. Give a running total update
	// }
	//ATTEMPT BUFIO ENDED HASN"T WOPKED"



//Pass each question into the getInput func. 
//User answer the question - check against score of QandAs slice.
// Updates score
// Moves to next question
// At end score is returned.