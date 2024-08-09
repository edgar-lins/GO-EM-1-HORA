package main

import "fmt"

type GameState struct {
	coins     int
	life      int
	blood     int
	character string
}

func (g GameState) GetCharacterAndBlood() (string, int) {
	return g.character, g.blood
}

func main() {
	state1 := GameState{
		blood:     13,
		character: "Mario",
		coins:     122,
		life:      2,
	}

	c, b := state1.GetCharacterAndBlood()
	fmt.Printf("character: %s blood: %d \n", c, b)

	state2 := GameState{
		blood:     63,
		character: "Princess",
		coins:     122,
		life:      4,
	}

	c2, b2 := state2.GetCharacterAndBlood()
	fmt.Printf("character: %s blood: %d", c2, b2)
}
