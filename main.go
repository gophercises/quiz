package main

import (
	"flag"
	"fmt"

	"github.com/fenriz07/quiz/cli"
	repositorycsv "github.com/fenriz07/quiz/repositories"
)

func main() {
	fmt.Println("Test")

	nameFile := flag.String("namefile", "problems.csv", "name to file evaluate")
	flag.Parse()

	csv := repositorycsv.ReadFile(*nameFile)

	cli.Show(csv)

}
