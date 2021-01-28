package game

// Game ..
type Game struct {
	Answer    int
	Opptunity int
}

// Guess ...
func (g *Game) Guess(num int) string {
	var s string = ""
	g.Opptunity--

	if g.Answer == num {
		s = "恭喜過關"
	} else {
		if g.Opptunity <= 0 {
			s = "失敗"
		} else {
			if g.Answer > num {
				s = "過小"
			} else {
				s = "過大"
			}
		}
	}

	return s
}
