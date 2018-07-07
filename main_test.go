package main

import (
	"testing"

	game "github.com/zkynet/cheatdice/Game"
)

func TestMain(t *testing.T) {

}

func TestDiceRoll3000000TimesMax6(t *testing.T) {
	p := game.Player{}
	p.DiceRolls = make(map[int]int)
	max := 6
	for i := 0; i < 3000000; i++ {
		p.ResetDice()
		p.RollDice(max)
		if p.DiceRolls[0] > max {
			t.Error("Expected a roll lower then", max, " but got :", p.DiceRolls[0])
			t.Fail()
		}
		if p.DiceRolls[1] > max {
			t.Error("Expected a roll lower then", max, " but got :", p.DiceRolls[1])
			t.Fail()
		}
	}
}

func BenchmarkDiceRoll(b *testing.B) {

	max := 6
	p := game.Player{}
	p.DiceRolls = make(map[int]int)

	for n := 0; n < b.N; n++ {
		p.RollDice(max)
	}
}
