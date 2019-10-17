package main

type Coordinates struct {
	x, y int
}

// Snake defintion
type Snake struct {
	Body      []Coordinates
	direction Direction
}

type Direction byte

const (
	Up Direction = iota
	Down
	Left
	Right
)

func NewSnake(c ...Coordinates) Snake {
	return Snake{c, Down}
}

func (s Snake) Head() Coordinates {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) Turn(d Direction) {
	s.direction = d
}

func (s Snake) Step() Coordinates {
	head := s.Head()

	switch s.direction {
	case Up:
		head.y--
	case Down:
		head.y++
	case Left:
		head.x--
	case Right:
		head.x++
	}

	return head
}

func (s *Snake) Move(head Coordinates) {
	s.Body = append(s.Body[1:], head)
}

func (s *Snake) Eat(head Coordinates) {
	s.Body = append(s.Body, head)
}

func (s Snake) Draw(f func(Coordinates)) {
	for _, c := range s.Body {
		f(c)
	}
}
