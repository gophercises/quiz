package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputBytes, err := ioutil.ReadFile("../../problems.csv")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(string(inputBytes))
}
