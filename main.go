package main

import (
	"fmt"
	"number-game/game"
)

func main() {
	fmt.Printf("hello, world\n")
	g := game.Game{
		Answer:    90,
		Opptunity: 3,
	}

	fmt.Println(g.Guess(9))
	fmt.Println(g.Guess(9))
	fmt.Println(g.Guess(9))

}
