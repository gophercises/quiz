package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type operation struct {
	formula string
	result  string
}

func createOperation(s string) (operation, error) {
	var o operation
	values := strings.Split(s, ",")
	if len(values) != 2 {
		return operation{}, errors.New("invalid csv line: " + s)
	}
	o.formula = values[0]
	o.result = values[1]
	return o, nil
}

func main() {
	csvFlag := flag.String("csv", "", "Custom .csv file")
	flag.Parse()
	var csvFilePath string
	if *csvFlag != "" {
		csvFilePath = *csvFlag
	} else {
		csvFilePath = "problems.csv"
	}
	csvFile, err := os.Open(csvFilePath)
	check(err)
	defer csvFile.Close()
	scanner := bufio.NewScanner(csvFile)
	scanner.Split(bufio.ScanLines)
	var total int
	var correct int
	for scanner.Scan() {
		o, err := createOperation(scanner.Text())
		check(err)
		fmt.Print(o.formula + " = ")
		var userInput string
		_, inputError := fmt.Scanf("%s", &userInput)
		if userInput == o.result && inputError == nil {
			correct++
		}
		total++
	}
	fmt.Println("TOTAL CORRECT ANSWERS: ", strconv.Itoa(correct))
	fmt.Println("TOTAL PROBLEMS: ", strconv.Itoa(total))
}
