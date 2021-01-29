package game

import (
	"testing"
)

// 單元測試重點
// 1. 程式碼覆蓋率
// 2. 邊界值

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

// TestGuessCodeFirst 先有code 再有testing case
func TestGuessCodeFirst(t *testing.T) {

	mGameCases := map[string]GameCase{}
	mGameCases["PassAt1st"] = GameCase{
		Game: Game{
			Answer:    90,
			Opptunity: 3,
		},
		Cases: []GuessCase{
			{
				Guess: 90,
				Want:  "恭喜過關",
			},
		},
	}

	mGameCases["PassAt4st"] = GameCase{
		Game: Game{
			Answer:    90,
			Opptunity: 3,
		},
		Cases: []GuessCase{
			{
				Guess: 91,
				Want:  "過大",
			},
			{
				Guess: 89,
				Want:  "過小",
			},
			{
				Guess: 1,
				Want:  "失敗",
			},
			{
				Guess: 90,
				Want:  "遊戲已結束",
			},
		},
	}

	for k, gameCase := range mGameCases {
		for _, guess := range gameCase.Cases {
			got := gameCase.Game.Guess(guess.Guess)
			if got != guess.Want {
				t.Errorf("Test for %v , Guess() = %v; want %v", k, got, guess.Want)
			}
		}
	}

}
