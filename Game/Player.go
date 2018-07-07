package game

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

type Player struct {
	Number     int
	Name       string
	Wins       int
	Rolls      int
	IsComputer bool // if true, rolls will happen automaticaly
	IsCheater  bool // if true, this player will get software assistance
	DiceRolls  map[int]int
}

func (p *Player) RollDice(min int, max int) {
	p.DiceRolls[0] = int(rollDice(int64(min), int64(max)))
	p.DiceRolls[1] = int(rollDice(int64(min), int64(max)))
}

func (p *Player) ResetDice() {
	p.DiceRolls[0] = 0
	p.DiceRolls[1] = 0
}

func rollDice(min int64, max int64) int64 {
	result, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Println("Can not generate random numbers.. globalGame halted. Error: ", err)
		os.Exit(1)
	}
	return result.Int64() + min
}