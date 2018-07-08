package game

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

type Player struct {
	Number      int
	Name        string
	Wins        int
	Rolls       int
	IsComputer  bool // if true, rolls will happen automaticaly
	IsCheater   bool // if true, this player will get software assistance
	CurrentDice map[int]int
}

func (p *Player) RollDice(max int) {
	p.Rolls++
	p.ResetDice()
	p.CurrentDice[0] = int(p.rollDice(int64(max)))
	p.CurrentDice[1] = int(p.rollDice(int64(max)))
}

func (p *Player) ResetDice() {
	p.CurrentDice[0] = 0
	p.CurrentDice[1] = 0
}

func (p *Player) rollDice(max int64) int64 {
	result, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Println("Can not generate random numbers.. globalGame halted. Error: ", err)
		os.Exit(1)
	}
	return result.Int64() + 1
}
