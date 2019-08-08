package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var csvfile = flag.String("csvfile", "problems.csv", "")

func main() {
	flag.Parse()

	file, err := os.Open(*csvfile)
	if err != nil {
		fmt.Println("Cannot find problems file.")
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	entries, err := reader.ReadAll()
	file.Close()
	if err != nil {
		fmt.Println("Malformed csv file.")
		os.Exit(1)
	}

	var correct int
	for _, e := range entries {
		q, a := e[0], e[1]
		fmt.Println(q)
		var u string
		fmt.Scanf("%s", &u)
		if u == a {
			correct++
		}
	}
	fmt.Printf("Correct answers: %d/%d\n", correct, len(entries))
}
