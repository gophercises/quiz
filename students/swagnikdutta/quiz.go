package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "../../problems.csv"
	file, err := os.Open(filepath)

	if err != nil { // handle error
		log.Fatal("Error reading csv\n", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	correct := 0
	input := bufio.NewScanner(os.Stdin)
	for idx, record := range records {
		fmt.Printf("Problem #%d: %s = ", idx+1, record[0])
		input.Scan()
		answer := input.Text() // string

		if answer != record[0] {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(records))
	file.Close()
}
