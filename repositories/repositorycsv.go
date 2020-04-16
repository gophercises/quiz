package repositorycsv

import (
	"encoding/csv"
	"io"
	"os"
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

	return csv

}
