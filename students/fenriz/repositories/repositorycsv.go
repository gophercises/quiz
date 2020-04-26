package repositories

import (
	"encoding/csv"
	"flag"
	"io"
	"math/rand"
	"os"
	"time"
)

type Row struct {
	question string
	answer   string
}

type Csv struct {
	records []Row
}

func (c Csv) GetRecords() []Row {
	return c.records
}

func (r Row) GetQuestion() string {
	return r.question
}

func (r Row) GetAnswer() string {
	return r.answer
}

func (c *Csv) Shuffle() {

	randCsv := flag.Bool("randcsv", false, "is rand quiz")

	flag.Parse()

	if *randCsv {
		rand.Seed(time.Now().UnixNano())

		for i := 1; i < len(c.records); i++ {
			r := rand.Intn(i + 1)
			if i != r {
				c.records[r], c.records[i] = c.records[i], c.records[r]
			}
		}
	}

}

func ReadFile(namefile string) Csv {

	csvfile, err := os.Open(namefile)

	if err != nil {
		println(err)
		panic(err)
	}

	r := csv.NewReader(csvfile)

	//records := [][]string{}

	csv := Csv{}

	for {

		// record is slice []string
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		row := Row{question: record[0], answer: record[1]}

		//fmt.Println(recordStruct)

		csv.records = append(csv.records, row)

		// slice record add in records slice
		//records = append(records, record)

	}

	csv.Shuffle()

	return csv

}
