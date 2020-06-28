package hero

// Hero is the protagonist.
type Hero interface {
	MoveUp()
	MoveDown()
	ReadyArrow()
	PullBow()
	ShootArrow()
}

// NewHero creates a hero.
func NewHero(mind Mind, body Body, quiver Quiver, bow Bow) Hero {
	return &hero{
		mind:   mind,
		body:   body,
		quiver: quiver,
		bow:    bow,
	}
}

type hero struct {
	mind   Mind
	body   Body
	quiver Quiver
	bow    Bow
}

func (hero *hero) MoveUp() {
	hero.body.MoveUp()
}

func (hero *hero) MoveDown() {
	hero.body.MoveDown()
}

func (hero *hero) ReadyArrow() {
	arrow, err := hero.quiver.Draw()
	if err != nil {
		hero.think("I'm out of arrows!")
	}

	hero.bow.ReadyArrow(arrow)
}

func (hero *hero) PullBow() {
	err := hero.bow.PullString()
	if err != nil {
		hero.think(err.Error())
	}
}

func (hero *hero) ShootArrow() {
	err := hero.bow.Release()
	if err != nil {
		hero.think("There are no arrows to shoot!")
	}
}

func (hero *hero) think(thought string) {
	hero.mind.ShowThought(thought)
}
