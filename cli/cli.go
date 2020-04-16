package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	repositorycsv "github.com/fenriz07/quiz/repositories"
)

func Show(csv repositorycsv.Csv) {

	resultUser := make(chan []string)

	go ask(csv, resultUser)

	calcResult(csv, resultUser)

}

func ask(csv repositorycsv.Csv, c chan<- []string) {

	reader := bufio.NewReader(os.Stdin)

	answers := []string{}

	for _, row := range csv.GetRecords() {

		fmt.Printf("Cuanto es el resultado de la siguente operación %v \n", row.GetQuestion())

		answer, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		answer = strings.Replace(answer, "\n", "", -1)

		answers = append(answers, answer)

	}

	c <- answers
}

func calcResult(csv repositorycsv.Csv, c <-chan []string) {
	var total int
	var corrects int
	var incorrects int

	answers := <-c

	for k, row := range csv.GetRecords() {
		if answers[k] == row.GetAnswer() {
			corrects++
		} else {
			incorrects++
		}

		total++
	}

	fmt.Printf("Numero de respuestas correctas %v \n", corrects)
	fmt.Printf("Numero de respuestas incorrectas %v \n", incorrects)
	fmt.Printf("Total de preguntas %v \n", total)

}
