package quiz

// Question represents a single quiz question with answer.
type Question struct {
	Question string
	Answer   string
}

// QuestionService implements all the methods related to a single quiz question.
type QuestionService interface {
	Ask() bool
}

// Quiz represents a quiz run with all the questions and answers.
type Quiz struct {
	questions []Question
	Results   []bool
}

// QuizService holds all the methods that can be applied to a Quiz.
type QuizService interface {
	Run()
	AddQuestion(Question)
}

func (q *Quiz) Run() {
	panic("not implemented")
}

func (q *Quiz) AddQuestion(question Question) {
	q.questions = append(q.questions, question)
}
