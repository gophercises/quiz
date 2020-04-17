package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	repositorycsv "github.com/fenriz07/quiz/repositories"
)

func Show(csv repositorycsv.Csv, limittime int) {

	resultUser := make(chan bool)

	timeoutCh := time.After(time.Duration(limittime) * time.Second)

	answers := []string{}

	go ask(csv, resultUser, &answers)

	select {
	case <-resultUser:
		fmt.Println("Gracias por terminar el test su resultado sera mostrado")
	case <-timeoutCh:
		fmt.Println("se acabo el tiempo su resultado sera mostrado")
	}

	calcResult(csv, answers)

}

func ask(csv repositorycsv.Csv, c chan<- bool, answers *[]string) {

	reader := bufio.NewReader(os.Stdin)

	for _, row := range csv.GetRecords() {

		fmt.Printf("Cuanto es el resultado de la siguente operación %v \n", row.GetQuestion())

		answer, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		answer = strings.Replace(answer, "\n", "", -1)

		*answers = append(*answers, answer)

	}

	c <- true
}

func calcResult(csv repositorycsv.Csv, answers []string) {
	var total int
	var corrects int
	var incorrects int

	lenanswers := len(answers)

	for k, row := range csv.GetRecords() {

		if k < lenanswers {
			if answers[k] == row.GetAnswer() {
				corrects++
			} else {
				incorrects++
			}
		} else {
			incorrects++
		}

		total++
	}

	fmt.Printf("Numero de respuestas correctas %v \n", corrects)
	fmt.Printf("Numero de respuestas incorrectas %v \n", incorrects)
	fmt.Printf("Total de preguntas %v \n", total)

}
