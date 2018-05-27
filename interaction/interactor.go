package interaction

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Interactor is used to fetch input from the terminal.
type Interactor struct {
	reader *bufio.Reader
	writer io.Writer
}

// Asker is used to fetch input from the terminal.
type Asker interface {
	Ask(string) string
	Notify(string)
}

// Ask asks an input to the user and expects a result.
func (i Interactor) Ask(question string) string {
	fmt.Print(question)
	text, _ := i.reader.ReadString('\n')
	return text
}

// Notify notifies the user of a message.
func (i Interactor) Notify(message string) {
	i.writer.Write([]byte(message))
}

// NewInteractor creates a interactor based asker.
func NewInteractor() Asker {
	r := bufio.NewReader(os.Stdin)
	return Interactor{reader: r, writer: os.Stdout}
}
