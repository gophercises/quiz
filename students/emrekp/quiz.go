package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filename := flag.String("file", "problems.csv", "Questions file")
	flag.Parse()

	csvF, csvErr := ioutil.ReadFile(*filename)
	if csvErr != nil {
		log.Fatal("Dosya bulunamadı veya bir sıkıntısı var. Detay: " + csvErr.Error())
	}

	probs, probErr := csv.NewReader(strings.NewReader(string(csvF))).ReadAll()
	if probErr != nil {
		log.Fatal("CSV düzgün formatlanmamış. Detay: " + probErr.Error())
	}

	var answer string
	var trues, total int

	for i, soru := range probs {
		fmt.Printf("%d. soru: %s = ", i+1, soru[0])
		fmt.Scan(&answer)
		if answer == soru[1] {
			trues++
			//log.Fatal(answer + " Cevap " + soru[1])
		}
		total++
	}

	fmt.Printf("True answers: %d\n", trues)
	fmt.Printf("Total questions: %d\n", total)
}
