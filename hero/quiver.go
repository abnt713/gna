package hero

import "errors"

// Quiver can hold arrows.
type Quiver interface {
	Count() uint
	Draw() (Arrow, error)
}

// NewQuiver creates a new quiver.
func NewQuiver(initialArrows uint, capacity uint) (Quiver, error) {
	if initialArrows > capacity {
		return nil, errors.New("Initial arrows cannot exceed the capacity")
	}

	return &theQuiver{
		arrowsCount: initialArrows,
		capacity:    capacity,
	}, nil
}

type theQuiver struct {
	arrowsCount uint
	capacity    uint
}

func (quiver *theQuiver) Count() uint {
	return quiver.arrowsCount
}

func (quiver *theQuiver) Draw() (Arrow, error) {
	if quiver.arrowsCount <= 0 {
		return nil, errors.New("Quiver is empty")
	}

	arrow := NewArrow()
	quiver.arrowsCount--
	return arrow, nil
}
