package hero

import "fmt"

// Body is the physical manifestation of a hero.
type Body interface {
	MoveUp()
	MoveDown()
}

// NewBody creates a new body.
func NewBody() Body {
	return &stdoutBody{}
}

type stdoutBody struct{}

func (body *stdoutBody) MoveUp() {
	fmt.Println("Moving up...")
}

func (body *stdoutBody) MoveDown() {
	fmt.Println("Moving down...")
}
