package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("problems.csv")
	if err != nil {
	}
	fileScanner := bufio.NewScanner(readFile)
	var correct = 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		var a = strings.Split(line, ",")

		fmt.Println(a[0])
		var digit int
		fmt.Scanf("%d", &digit)
		var a5, _ = strconv.Atoi(a[1])
		if digit == a5 {
			correct++
		}
	}
	if correct > 11 {
		var cor = strconv.Itoa(correct)
		fmt.Println("You scored " + cor + " out of 12 😄")
	} else {
		var cor = strconv.Itoa(correct)
		fmt.Println("You scored " + cor + " out of 12 at least you tried... 😭")
	}
}
