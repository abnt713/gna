package hero

import "fmt"

// Arrow defines an arrow :D
type Arrow interface {
	GetLaunchedByStrenght(strength int)
}

// NewArrow creates an arrow.
func NewArrow() Arrow {
	return &theArrow{}
}

type theArrow struct{}

func (*theArrow) GetLaunchedByStrenght(strength int) {
	fmt.Printf("[ARROW] Getting launched with a strength of %v\n", strength)
}
