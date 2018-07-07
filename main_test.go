package main

import (
	"os"
	"strconv"
	"testing"

	game "github.com/zkynet/cheatdice/Game"
)

func TestMain(t *testing.T) {
	loadGame()
}

func loadGame() {
	computer0 := game.Player{
		Name:       "AlphaDice0",
		Wins:       0,
		Rolls:      0,
		IsComputer: true,
		IsCheater:  true,
		DiceRolls:  make(map[int]int),
	}

	computer1 := game.Player{
		Name:       "AlphaDice1",
		Wins:       0,
		Rolls:      0,
		IsComputer: true,
		IsCheater:  false,
		DiceRolls:  make(map[int]int),
	}
	globalGame = &game.Game{}
	globalGame.InitGame()
	globalGame.Players[0] = &computer0
	globalGame.Players[1] = &computer1
	globalGame.CheatCounter = 0
	globalGame.FirstCheatRound = 1
	globalGame.WinningPercent = 0.7
	globalGame.CheatsInARow = 2
	globalGame.NumberOfCheatMethods = 3
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

func TestWinningPercentage100Rounds70WinRating(t *testing.T) {

	f, err := os.OpenFile("100-round-roll-log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < 100; i++ {
		globalGame.Round = globalGame.Round + 1
		globalGame.ResetAllDice()
		rollDiceForPlayer(globalGame.CurrentRoller)
		globalGame.SwitchPlayers()
		rollDiceForPlayer(globalGame.CurrentRoller)
		_, _ = globalGame.FindRoundWinner()
		globalGame.CalculateWinRatings()

		text := strconv.Itoa(globalGame.Players[0].DiceRolls[0]) + strconv.Itoa(globalGame.Players[0].DiceRolls[1]) + "-" + strconv.Itoa(globalGame.Players[1].DiceRolls[0]) + strconv.Itoa(globalGame.Players[1].DiceRolls[1]) + "\n"
		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
	}

	if globalGame.WinRatings[0] < 0.6 || globalGame.WinRatings[0] > 0.8 {
		t.Error("Win rating was:", globalGame.WinRatings[0], "wanted a value between 0.6 and 0.8")
		t.Fail()
	}

	//fmt.Println("Win rating:", globalGame.WinRatings[0])

}

func TestWinningPercentage1000000Rounds70WinRating(t *testing.T) {

	f, err := os.OpenFile("1000000-round-roll-log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i := 0; i < 1000000; i++ {
		globalGame.Round = globalGame.Round + 1
		globalGame.ResetAllDice()
		rollDiceForPlayer(globalGame.CurrentRoller)
		globalGame.SwitchPlayers()
		rollDiceForPlayer(globalGame.CurrentRoller)
		_, _ = globalGame.FindRoundWinner()
		globalGame.CalculateWinRatings()

		text := strconv.Itoa(globalGame.Players[0].DiceRolls[0]) + strconv.Itoa(globalGame.Players[0].DiceRolls[1]) + "-" + strconv.Itoa(globalGame.Players[1].DiceRolls[0]) + strconv.Itoa(globalGame.Players[1].DiceRolls[1]) + "\n"
		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
	}

	if globalGame.WinRatings[0] < 0.6 || globalGame.WinRatings[0] > 0.8 {
		t.Error("Win rating was:", globalGame.WinRatings[0], "wanted a value between 0.6 and 0.8")
		t.Fail()
	}

	//fmt.Println("Win rating:", globalGame.WinRatings[0])

}

func BenchmarkDiceRoll(b *testing.B) {

	max := 6
	p := game.Player{}
	p.DiceRolls = make(map[int]int)

	for n := 0; n < b.N; n++ {
		p.RollDice(max)
	}
}
