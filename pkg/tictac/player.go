package tictac

type Player struct {
	Card   *Card
	Player int
	Round  int
	s      string
}

func NewPlayer(id int, s string) *Player {
	return &Player{
		s: s,
	}
}
