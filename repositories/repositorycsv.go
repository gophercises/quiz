package repositorycsv

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadFile(namefile string) [][]string {

	csvfile, err := os.Open(namefile)

	if err != nil {
		println(err)
		panic(err)
	}

	r := csv.NewReader(csvfile)

	records := [][]string{}

	for {

		// record is slice []string
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		// slice record add in records slice
		records = append(records, record)

	}

	return records

}
