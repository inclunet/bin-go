package tictac

import (
	"net/http"

	"github.com/inclunet/bin-go/pkg/server"
)

type TicTac struct {
	Rounds []*Round
}

func (t *TicTac) AddRoundsHandler(r *http.Request) (*server.Response, error) {
	Round := NewRound(len(t.Rounds), server.GetURLParamHasInt(r, "type"))

	Round.AddPlayer("O")

	t.Rounds = append(t.Rounds, Round)

	return server.NewResponse(Round.Card)
}
