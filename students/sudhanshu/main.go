package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Data struct {
	question string `json:"question"`
	answer string  `json:"answer"`
}


func ReadCSV(csvFile string) []Data {
	var dataset []Data
	//load csv
	f, _ := os.Open(csvFile)

	//create a reader
	r := csv.NewReader(f)

	for{
		record, err := r.Read()

		if err == io.EOF {
			break;
		}
		if err != nil{
			panic("Error occurred while reading")
		}
		dataset = append(dataset, Data{record[0], record[1]})
	}
	return dataset
}


func main()  {
 dataset := ReadCSV("../../problems.csv")
 var score int = 0
 var user_answer string


 for i, data := range dataset{
 	fmt.Print("Problem #",i+1,": ",data.question," = ")
	 _, _ = fmt.Scanln(&user_answer)

	 // trim string
	 //user_answer = strings.Trim(user_answer, " ")

	 if user_answer == data.answer{
	 	score  +=1
	 }

 }

println("Your score is :", score,"/",len(dataset))

}