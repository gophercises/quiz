package main

import (
  "fmt"
  "flag"
  "bufio"
  "time"
  "os"
  "strings"
  "math/rand"
)

type Pair struct {
  question string
  answer string
}

func main(){
  // TODO: CLI parameters done.
  var csvlocation *string
  var timeduration *int
  var shuffle *string
  csvlocation = flag.String("csv", "problems.csv", "absolute location of question file")
  timeduration = flag.Int("time", 30, "timelength")
  shuffle = flag.String("shuffle", "n", "if you want shuffling, enter 'y'")
  flag.Parse()

  // TODO: open file
  csvFile, err := os.Open(*csvlocation)

  if err != nil {
    panic(err)
  }
  fileReader := bufio.NewScanner(csvFile)

  // TODO: open question array
  var query []Pair

  // TODO: read from file
  for fileReader.Scan(){
    line := strings.Split(fileReader.Text(), ",")
    q := strings.TrimSpace(line[0])
    a := strings.TrimSpace(line[1])

    query = append(query, Pair{question: q, answer: a})
  }

  // TODO: go routine
  var trues int
  var totals int = len(query)

  // TODO: Shuffling
  if strings.EqualFold(*shuffle, "y") {
    shuffled := make([]Pair, len(query))
    rand.Seed(time.Now().UTC().Unix())
    permutation := rand.Perm(len(query))

    for i, v := range permutation {
      shuffled[v] = query[i]
    }
    go func(){
      for i, pair := range shuffled {
        fmt.Printf("Question %d: %s\n", (i+1), pair.question)

        var input string
        fmt.Scan(&input)

        if pair.answer == input {
          trues++
        }
      }
    }()
  } else {
    go func(){
      for i, pair := range query {
        fmt.Printf("Question %d: %s\n", (i+1), pair.question)

        var input string
        fmt.Scan(&input)

        if strings.EqualFold(pair.answer, input) {
          trues++
        }
      }
      }()
  }

// TODO: Timer sequence
  timer := time.NewTimer(time.Second * time.Duration(*timeduration))

  <-timer.C

  // TODO: Print results
  fmt.Printf("Right answer: %d\nTotal answer: %d\n", trues, totals)
}
