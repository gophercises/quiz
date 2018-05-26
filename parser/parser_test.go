package parser

import (
	"testing"

	"github.com/fedepaol/quiz/quiz"
)

var qtests = []struct {
	row []string
	q   quiz.Question
}{
	{[]string{"1+2", "3"}, quiz.Question{"1+2", "3"}},
	{[]string{"what 2+2, sir?", "4"}, quiz.Question{"what 2+2, sir?", "4"}},
}

func TestParsingQuestion(t *testing.T) {
	for _, tt := range qtests {
		q, e := parseQuestion(tt.row)
		if e != nil {
			t.Errorf("Error for %q", tt.row)
		}

		if q != tt.q {
			t.Errorf("got %q, want %q", q, tt.q)
		}
	}
}
