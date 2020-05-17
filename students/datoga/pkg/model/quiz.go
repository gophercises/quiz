package model

// QuizItem is a question in the quiz
type QuizItem struct {
	Question string
	Solution string
}

// Quiz is a limited set of QuizItem
type Quiz []QuizItem
