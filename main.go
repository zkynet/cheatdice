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
	globalGame.FirstCheatRound = 1
	globalGame.WinningPercent = 0.7
	globalGame.CheatsInARow = 2
	globalGame.NumberOfCheatMethods = 3

	fmt.Println("Press r to roll the dice or press s to stop the globalGame")

	startGame()
}

func startGame() {
NEXTROUND:
	globalGame.Round = globalGame.Round + 1
	fmt.Println("====================================================")
	fmt.Println("======== GAME ROUND: ", globalGame.Round, "=========")
	fmt.Println("====================================================")

	globalGame.ResetAllDice()
	globalGame.AskPlayerToRoll()
	// all printing methods are customized outside of the game logic
	// this allows for more customization
	printDiceRollForPlayer(globalGame.CurrentRoller)
	globalGame.SwitchPlayers()
	globalGame.AskPlayerToRoll()
	printDiceRollForPlayer(globalGame.CurrentRoller)
	// the _ variable is the winning player struct, if we wanted to use the game in another program
	// we would skip printing and just get the winnig player.
	_, message := globalGame.FindRoundWinner()
	fmt.Println(message)
	globalGame.CalculateWinRatings()
	printStandings()

	fmt.Println()
	fmt.Println()
	fmt.Println()
	goto NEXTROUND
}

func printDiceRollForPlayer(number int) {
	fmt.Println("Dice Roll: ", globalGame.Players[number].Name, " (", globalGame.Players[number].DiceRolls[0], "/", globalGame.Players[number].DiceRolls[1], ")")
}

func printStandings() {
	fmt.Println("=== Current standings ===")
	fmt.Println(globalGame.Players[0].Name, " has ", globalGame.Players[0].Wins, "wins witha win rating of", globalGame.WinRatings[0])
	fmt.Println(globalGame.Players[1].Name, " has ", globalGame.Players[1].Wins, "wins witha win rating of", globalGame.WinRatings[1])
}
