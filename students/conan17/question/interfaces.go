package question

type Quizzes interface {
	Next() error
	Launch()
	Summary()
	Size() int
	Remain() int
	Random(bool)
}

type Quiz interface {
	Ask()
	Answer() bool
}
