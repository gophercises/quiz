package interaction

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Interactor is used to fetch display questions and retrieve replies.
type Terminal struct {
	reader *bufio.Reader
	writer io.Writer
}

// Asker is used to fetch input from the terminal.
type Asker interface {
	Ask(string) string
	Notify(string)
}

// Ask asks an input to the user and expects a result.
func (a Terminal) Ask(question string) string {
	fmt.Printf("%s = ", question)
	text, _ := a.reader.ReadString('\n')
	return text
}

// Notify notifies the user of a message.
func (a Terminal) Notify(message string) {
	a.writer.Write([]byte(message))
}

// NewAsker creates a interactor based asker.
func NewAsker() Asker {
	r := bufio.NewReader(os.Stdin)
	return Terminal{reader: r, writer: os.Stdout}
}
