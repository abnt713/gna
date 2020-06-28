package hero

import "fmt"

// Mind display thoughts.
type Mind interface {
	ShowThought(thought string)
}

// NewStdoutMind creates a mind which displays thoughts throught stdout.
func NewStdoutMind() Mind {
	return &stdoutMind{}
}

type stdoutMind struct{}

func (stdoutMind *stdoutMind) ShowThought(thought string) {
	fmt.Println(thought)
}
