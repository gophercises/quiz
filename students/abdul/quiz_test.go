package quiz

import (
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

func testEachQuestion(t *testing.T) {
	timer := time.NewTimer(time.Duration(2) * time.Second).C
	done := make(chan string)
	var quest Question
	quest.question = "1+1"
	quest.answer = "2"
	var ans int
	var err error
	allDone := make(chan bool)
	go func() {
		ans, err = eachQuestion(quest.question, quest.answer, timer, done)
		allDone <- true
	}()
	done <- "2"

	<-allDone
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, ans, 1)
}

func testReadCSV(t *testing.T) {
	str := "1+1,2\n2+1,3\n9+9,18\n"
	quest, err := readCSV(strings.NewReader(str))
	if err != nil {
		t.Error(err)
	}
	var que [3]Question
	que[0].answer = "2"
	que[1].answer = "3"
	que[2].answer = "18"
	que[0].question = "1+1"
	que[1].question = "2+1"
	que[2].question = "9+9"

	assert.Equal(t, que[0], quest[0])
	assert.Equal(t, que[1], quest[1])
	assert.Equal(t, que[2], quest[2])

}

func TestEachQuestion(t *testing.T) {
	t.Run("test eachQuestion", testEachQuestion)
}

func TestReadCSV(t *testing.T) {
	t.Run("test ReadCSV", testReadCSV)
}
