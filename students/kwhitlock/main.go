package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func sendQuestion (text string) {
	fmt.Println("New question: ")
	fmt.Println(text)
}

func main(){
	fmt.Println("Running Quiz")

	file, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("Could not open file")
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
	for scanner.Scan() {
		sendQuestion(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR")
		log.Fatal(err)
	}
	fmt.Println("End of quiz")



}