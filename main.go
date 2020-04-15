package main

import (
	"fmt"

	repositorycsv "github.com/fenriz07/quiz/repositories"
)

func main() {
	fmt.Println("Probando")

	records := repositorycsv.ReadFile("problems.csv")

	fmt.Println(records)
}
