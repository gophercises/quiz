package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	problemsFilename := flag.String("c", "problems.csv", "The file containing the problems")
	flag.Parse()

	file, err := os.Open(*problemsFilename)
	if err != nil {
		fmt.Printf("Failed to open %s\n", *problemsFilename)
		os.Exit(1)
	}

	csv := csv.NewReader(file)
	lines, err := csv.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read %s\n", *problemsFilename)
		os.Exit(1)
	}

	var text string
	correct := 0

	for _, line := range lines {
		question := strings.TrimSpace(line[0])
		answer := strings.TrimSpace(line[1])
		fmt.Printf("%s? ", question)
		fmt.Scanf("%s\n", &text)

		if text == answer {
			correct++
		}
	}

	fmt.Printf("You got %d correct output of %d\n", correct, len(lines))
}
