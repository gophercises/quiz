package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	csvReaderRow()

}

func csvReaderRow() {
	recordFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("an error has encoutered opening file", err)
		return
	}

	reader := csv.NewReader(recordFile)

	header, err := reader.Read()
	if err != nil {
		fmt.Println("an error encountered reading file", err)
		return
	}
	fmt.Printf("Headers : %v  \n", header)

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("an error encountered recording answer", err)
			return
		}

		fmt.Printf("Row %d : %v \n", i, record)
	}

}
