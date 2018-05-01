package game

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"quiz/problem"
	"strings"
	"time"
)

// Engine type contains all the data needed for the game to run
type Engine struct {
	correctAnswers map[bool]int
	questions      []*problem.Question
	quizRandomness bool
}

// NewEngine creates a new game engine which controls the game runtime.
// Returns a pointer to the created Engine type variable.
func NewEngine() *Engine {
	engine := Engine{
		correctAnswers: map[bool]int{},
		quizRandomness: false,
	}

	return &engine
}

// SetQuizProblems loads new quiz questions from the given csvFilePath.
// Any previous loaded problems will be removed prior to loading new ones.
// Returns error if e.g. there were problems with parsing the file specified
// in csvFilePath.
func (e *Engine) SetQuizProblems(csvFilePath string) error {
	e.questions = []*problem.Question{}
	questionAnswers, err := parseProblemsFrom(csvFilePath)
	if err != nil {
		return err
	}

	for _, qa := range questionAnswers {
		e.questions = append(e.questions,
			problem.NewQuestion(strings.TrimSpace(qa[0]), strings.TrimSpace(qa[1])))
	}

	return nil
}

func parseProblemsFrom(csvFilePath string) (questionsAnswers [][]string, err error) {
	problemsCsvFile, err := os.Open(csvFilePath)
	if err != nil {
		pathErr := err.(*os.PathError)
		return questionsAnswers, fmt.Errorf("Failed to perform %v on file %v: %v", pathErr.Op, pathErr.Path, pathErr.Error())
	}

	csvReader := csv.NewReader(problemsCsvFile)
	questionsAnswers, err = csvReader.ReadAll()
	if err != nil {
		return questionsAnswers, fmt.Errorf("Failed to parse problems: %v", err)
	}

	return questionsAnswers, nil
}

// SetQuizRandomness sets whether the order of the quiz questions
// will be randomized on each game round or not.
// Quiz randomness is false by default.
func (e *Engine) SetQuizRandomness(random bool) {
	e.quizRandomness = random
}

// StartGame starts and runs one round of quiz questions. As a prerequisite, it's required to
// load the quiz problems first (see SetQuizProblems() method).
// Returns error if there were problems reading input from the user.
// To print results after the game is run, use PrintGameResults() method.
func (e *Engine) StartGame() (err error) {
	if len(e.questions) == 0 {
		return fmt.Errorf("Quiz questions not loaded. Use Engine.SetQuizProblems() to load question set")
	}

	// initialize new game round
	fmt.Println("Starting quiz...")
	e.correctAnswers[false] = 0
	e.correctAnswers[true] = 0

	// ask questions to user
	return e.answerQuestions()
}

func (e *Engine) answerQuestions() error {
	// if random order of questions wanted, randomize indices for iteration
	// otherwise set indices for iteration in order
	var iterationIndices []int
	if e.quizRandomness {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		iterationIndices = r.Perm(len(e.questions))
	} else {
		iterationIndices = make([]int, len(e.questions))
		for i := range iterationIndices {
			iterationIndices[i] = i
		}
	}

	// ask user to give answer for each of the questions
	// and store which are correct/incorrect in e.correctAnswers
	reader := bufio.NewReader(os.Stdin)
	for _, iterIndex := range iterationIndices {
		q := e.questions[iterIndex]
		fmt.Printf("%v: ", q.Question())
		answer, err := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if err != nil {
			return fmt.Errorf("Failed to read from stdin: %v", err)
		}
		e.correctAnswers[q.Give(answer)]++
	}

	return nil
}

// PrintGameResults prints the results of the game last played
func (e *Engine) PrintGameResults() {
	fmt.Println("Correct answers:", e.correctAnswers[true])
	fmt.Println("Incorrect answers:", e.correctAnswers[false])
}
