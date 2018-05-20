package problem

// Problem represents a single question answer pair
type Problem struct {
	question string
	answer   string
}

// New creates a Problem from a provided CSV record
func New(record []string) Problem {
	return Problem{
		question: record[0],
		answer:   record[1],
	}
}
