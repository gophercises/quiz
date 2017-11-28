package reader

import (
	"log"
	"os"
	"testing"
)

func TestReadCsv(t *testing.T) {
	name := "test.csv"
	f := createFile(name)

	dat := ReadCsv(name)
	if len(dat) != 2 {
		log.Println("Map size differs than expected in readfile")
		t.Fail()
	}

	if dat["kOne"] != "valOne" && dat["kTwo"] != "valTwo" {
		log.Println("Fail while reading file content")
		t.Fail()
	}
	f.Close()
}
func createFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	f.WriteString("kOne,valOne,")
	f.WriteString("\n")
	f.WriteString("kTwo,valTwo,")
	f.WriteString("\n")

	return f
}
