package hero

import "errors"

// Bow can shoot arrows.
type Bow interface {
	ReadyArrow(arrow Arrow) error
	PullString() error
	Release() error
}

// NewBow creates a bow.
func NewBow() Bow {
	return &theBow{
		state: idle,
	}
}

type bowState int

const (
	idle bowState = iota
	loaded
	tense
)

type theBow struct {
	state bowState

	currArrow Arrow
}

func (bow *theBow) ReadyArrow(arrow Arrow) error {
	if bow.state != idle {
		return errors.New("There is an arrow already loaded")
	}
	bow.currArrow = arrow
	bow.state = loaded
	return nil
}

func (bow *theBow) PullString() error {
	if bow.state != loaded {
		return errors.New("The bow isn't loaded")
	}

	bow.state = tense
	return nil
}

func (bow *theBow) Release() error {
	if bow.state != tense {
		return errors.New("The bow isn't tense")
	}

	bow.currArrow.GetLaunchedByStrenght(10)
	bow.currArrow = nil

	bow.state = idle
	return nil
}
