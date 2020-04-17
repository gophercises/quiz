package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fenriz07/quiz/cli"
	repositorycsv "github.com/fenriz07/quiz/repositories"
)

func main() {
	fmt.Println("Iniciando el tests.")

	nameFile := flag.String("namefile", "problems.csv", "name to file evaluate")
	limitTime := flag.Int("limittime", 30, "time limit for quiz")

	flag.Parse()

	welcome(*limitTime)

	csv := repositorycsv.ReadFile(*nameFile)

	cli.Show(csv, *limitTime)

}

func welcome(limittime int) {

	fmt.Printf("Su test tiene una duración de %v segundos \n", limittime)
	fmt.Println("Presione enter para iniciar")

	reader := bufio.NewReader(os.Stdin)

	_, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
}
