package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("./problems.csv")
	r := csv.NewReader(io.Reader(file))

	data, _ := r.ReadAll()

	correct := 0
	total := len(data)
	for _, problem := range data {
		question := problem[0]
		answer := problem[1]
		fmt.Printf("%s: ", question)

		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess)
		if strings.Compare(answer, guess) == 0 {
			correct++
		}
	}

	fmt.Printf("You answered %d of %d correct", correct, total)
}
