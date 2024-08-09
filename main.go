package main

import "fmt"

type MarioGame struct {
	coins     int
	life      int
	blood     int
	character string
}

func (g MarioGame) GetCharacterAndBlood() (string, int) {
	return g.character, g.blood
}

type KirbyGame struct {
	character string
	blood     int
	power     string
}

func (k KirbyGame) GetCharacterAndBlood() (string, int) {
	return k.character, k.blood
}

type GameGetter interface {
	GetCharacterAndBlood() (string, int)
}

func main() {
	var marioState GameGetter = MarioGame{
		blood:     13,
		character: "Mario",
		coins:     122,
		life:      2,
	}

	c, b := marioState.GetCharacterAndBlood()
	fmt.Println(("===== MARIO"))
	fmt.Printf("character: %s blood: %d \n", c, b)

	var kirbuyState GameGetter = KirbyGame{
		character: "kirby",
		blood:     123,
	}

	ch, bl := kirbuyState.GetCharacterAndBlood()
	fmt.Println(("===== KIRBY"))
	fmt.Printf("character: %s blood: %d \n", ch, bl)
}
