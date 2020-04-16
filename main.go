package main

import (
	"fmt"

	"github.com/fenriz07/quiz/cli"
	repositorycsv "github.com/fenriz07/quiz/repositories"
)

func main() {
	fmt.Println("Test")

	csv := repositorycsv.ReadFile("problems.csv")

	cli.Show(csv)

}
