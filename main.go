package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	game "github.com/zkynet/cheatdice/Game"
)

// global globalGame struct .. etc ..
var globalGame *game.Game

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("===================================")
	fmt.Println("How many human players are playing ?")
	fmt.Println("===================================")
	playerCount, _, err := reader.ReadLine()
	playerCountAsInt, err := strconv.Atoi(string(playerCount))
	if err != nil {
		fmt.Println("Could not parse the player count, err:", err)
		os.Exit(1)
	}

	fmt.Println("===================================")
	fmt.Println("How many computer players are playing ?")
	fmt.Println("===================================")
	playerCountComputer, _, err := reader.ReadLine()
	playerCountComputerAsInt, err := strconv.Atoi(string(playerCountComputer))
	if err != nil {
		fmt.Println("Could not parse the player count, err:", err)
		os.Exit(1)
	}

	globalGame = &game.Game{}
	globalGame.InitGame()

	// cheat settings
	globalGame.CheatCounter = 0
	globalGame.FirstCheatRound = 1
	globalGame.WinningPercent = 0.7
	globalGame.CheatsInARow = 2
	globalGame.NumberOfCheatMethods = 3

	for i := 0; i < playerCountAsInt; i++ {
		fmt.Println("Enter Human Player #", i+1, " name:")
		name, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Could not parse the player name, err:", err)
			os.Exit(1)
		}
		player := game.Player{
			Name:        string(name),
			Wins:        0,
			Rolls:       0,
			IsComputer:  false,
			IsCheater:   true,
			CurrentDice: make(map[int]int),
		}

		if string(name) == "AlphaDice" {
			player.IsCheater = true
		}

		globalGame.Players[i] = &player
	}

	humanPlayercount := len(globalGame.Players)
	for i := humanPlayercount; i < (playerCountComputerAsInt + humanPlayercount); i++ {

		player := game.Player{
			Name:        "Computer-" + strconv.Itoa((i - humanPlayercount)),
			Wins:        0,
			Rolls:       0,
			IsComputer:  true,
			IsCheater:   false,
			CurrentDice: make(map[int]int),
		}
		globalGame.Players[i] = &player
	}

	fmt.Println("Press r to roll the dice or press s to stop the globalGame")

	startGame()
}

func startGame() {
NEXTROUND:
	// bump the game round
	globalGame.StartRound()
	globalGame.ResetAllDice()

	fmt.Println("====================================================")
	fmt.Println("======== GAME ROUND: ", globalGame.Round, "=========")
	fmt.Println("====================================================")

	// roll for all players
	for i := 0; i < len(globalGame.Players); i++ {
		globalGame.AskPlayerToRoll()
		printDiceRollForPlayer(globalGame.CurrentRoller)
		globalGame.SwitchPlayers()

	}

	globalGame.SwitchStartingPlayer()
	_, message := globalGame.FindRoundWinner()
	globalGame.CalculateWinRatings()
	fmt.Println(message)
	printStandings()

	fmt.Println()
	fmt.Println()
	goto NEXTROUND
}

func printDiceRollForPlayer(number int) {
	fmt.Println("Dice Roll: ", globalGame.Players[number].Name, " (", globalGame.Players[number].CurrentDice[0], "/", globalGame.Players[number].CurrentDice[1], ")")
}

func printStandings() {
	fmt.Println("=== Current standings ===")
	for i, p := range globalGame.Players {
		fmt.Println(p.Name, "has", p.Wins, "wins with a win rating of", globalGame.WinRatings[i])
	}
	fmt.Println("And there have been", globalGame.TieCount, "ties")
}
