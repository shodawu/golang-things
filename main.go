package main

import (
	"fmt"
	"log"
	"number-game/game"
	"number-game/hubs"
)

func main() {
	fmt.Printf("hello, world\n")

	lstSample := hubs.Controller{}
	lstSample.Hubs = game.LoadSamples()
	for _, cmd := range lstSample.Hubs {
		ok := cmd.Exec()
		log.Println("testing passed: ", ok)
	}

}
