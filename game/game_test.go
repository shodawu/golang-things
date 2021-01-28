package game

import (
	"testing"
)

type GameCase struct {
	Game  Game
	Cases []GuessCase
}

type GuessCase struct {
	Guess int
	Want  string
}

func TestGuess(t *testing.T) {

	tGame := GameCase{
		Game: Game{
			Answer:    90,
			Opptunity: 3,
		},
		Cases: []GuessCase{
			{
				Guess: 8,
				Want:  "過小",
			},
			{
				Guess: 99,
				Want:  "過大",
			},
			{
				Guess: 99,
				Want:  "失敗",
			},
		},
	}

	for _, guess := range tGame.Cases {
		got := tGame.Game.Guess(guess.Guess)
		if got != guess.Want {
			t.Errorf("Guess() = %v; want %v", got, guess.Want)
		}
	}

}
