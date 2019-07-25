package quiz

import (
	"testing"
	"time"
)

var qq = []Question{
	Question{"1+2", "3"},
	Question{"2+2", "4"},
}

var rr = map[string]string{
	"1+2": "3",
	"2+2": "2",
}

type MockAsker struct {
}

func (m MockAsker) Ask(q string) string {
	return rr[q]
}

func (m MockAsker) Notify(string) {
	panic("not implemented")
}

func TestQuiz(t *testing.T) {
	reschan := make(chan Result)
	q := Quiz{questions: qq, Asker: MockAsker{}}
	timeout := make(chan time.Time)
	go func() {
		reschan <- q.Run(timeout)
	}()

	res := <-reschan
	if res.questionsasked != 2 {
		t.Errorf("expecting 2 questions")
	}
}

type MockAskerTimeout struct {
	timeout chan time.Time
}

func (m MockAskerTimeout) Ask(q string) string {
	if q == "2+2" {
		m.timeout <- time.Now()
	}
	return rr[q]
}

func (m MockAskerTimeout) Notify(string) {
	panic("not implemented")
}

func TestTimeout(t *testing.T) {
	reschan := make(chan Result)

	a := MockAskerTimeout{timeout: make(chan time.Time)}
	q := Quiz{questions: qq, Asker: a}
	go func() {
		reschan <- q.Run(a.timeout)
	}()

	res := <-reschan
	if res.questionsasked != 1 {
		t.Errorf("expecting 1 questions")
	}
}
