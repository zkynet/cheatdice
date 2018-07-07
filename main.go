package main

import (
	"bufio"
	"fmt"
	"os"

	game "github.com/zkynet/cheatdice/Game"
)

// global globalGame struct .. etc ..
var globalGame *game.Game

func main() {
	fmt.Println("Human player name:")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Could not parse your name, err:", err)
		os.Exit(1)
	}

	computer := game.Player{
		Name:       "AlphaDice:)",
		Wins:       0,
		Rolls:      0,
		IsComputer: true,
		IsCheater:  true,
		DiceRolls:  make(map[int]int),
	}

	human := game.Player{
		Name:       string(name),
		Wins:       0,
		Rolls:      0,
		IsComputer: false,
		IsCheater:  false,
		DiceRolls:  make(map[int]int),
	}

	globalGame = &game.Game{}
	globalGame.InitGame()
	globalGame.Players[0] = &computer
	globalGame.Players[1] = &human

	// cheat settings
	globalGame.CheatCounter = 0
	globalGame.CheatOnRound = 10
	globalGame.WinningPercent = 0.7

	// globalGame loop
	// 1. read input
	// 1.1 if it is the computers turn we need to analyze score and determine if we need to cheat
	// -- if we cheat we can only do it when the player goes first
	// -- needs to win X percent of globalGames
	// -- can never cheat three in a row
	// 2. roll or stop
	// 3.

	fmt.Println(globalGame)
	fmt.Println("Press r to roll the dice or press s to stop the globalGame")

LOOP:
	globalGame.Round = globalGame.Round + 1
	fmt.Println("====================================================")
	fmt.Println("======== GAME ROUND: ", globalGame.Round, "=========")
	fmt.Println("====================================================")
	nextRound()
	goto LOOP
}

func nextRound() {
	globalGame.ResetAllDice()
	rollDiceForPlayer(globalGame.CurrentRoller)
	printDiceRollForPlayer(globalGame.CurrentRoller)
	globalGame.SwitchPlayers()
	rollDiceForPlayer(globalGame.CurrentRoller)
	printDiceRollForPlayer(globalGame.CurrentRoller)
	_, message := globalGame.FindRoundWinner()
	fmt.Println(message)
	globalGame.CalculateWinRatings()
	printStandings()
}

func userInputPrompt() {
	// do not move outside of loop, this will cause new line(enter) to be read as the next character
	reader := bufio.NewReader(os.Stdin)

PROMPTLOOP:
	switch readChar(reader) {
	case 'R', 'r':
		return
	case 'S', 's':
		os.Exit(1)
	default:
		fmt.Println("You pressed something other then r or s, please try again..")
		goto PROMPTLOOP
	}

}

func rollDiceForPlayer(number int) {

	if globalGame.Players[number].IsComputer {
		if globalGame.Cheat() {
			// if the player is a computer and it's cheating then we don't need to roll the dice.
			return
		}
	} else {
		userInputPrompt()
	}

	globalGame.Players[number].ResetDice()
	globalGame.Players[number].RollDice(globalGame.Dice.Max)
}

func readChar(reader *bufio.Reader) rune {
READ:
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println("An unexcpected error occurred: ", err)
		goto READ
	}

	return char
}

func printDiceRollForPlayer(number int) {
	fmt.Println("Dice Roll: ", globalGame.Players[number].Name, " (", globalGame.Players[number].DiceRolls[0], "/", globalGame.Players[number].DiceRolls[1], ")")
}

func printStandings() {

	fmt.Println("=== Current standings ===")
	fmt.Println(globalGame.Players[0].Name, " has ", globalGame.Players[0].Wins, "wins witha win rating of", globalGame.WinRatings[0])
	fmt.Println(globalGame.Players[1].Name, " has ", globalGame.Players[1].Wins, "wins witha win rating of", globalGame.WinRatings[1])

}
