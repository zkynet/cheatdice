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

func getPlayerCount() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("==========================================================")
	fmt.Println("How many human players are playing ?")
	fmt.Println("==========================================================")
	playerCount, _, err := reader.ReadLine()
	playerCountAsInt, err := strconv.Atoi(string(playerCount))
	if err != nil {
		fmt.Println("Could not parse the player count, err:", err)
		os.Exit(1)
	}

	fmt.Println("==========================================================")
	fmt.Println("How many computer players are playing ?")
	fmt.Println("==========================================================")
	playerCountComputer, _, err := reader.ReadLine()
	playerCountComputerAsInt, err := strconv.Atoi(string(playerCountComputer))
	if err != nil {
		fmt.Println("Could not parse the player count, err:", err)
		os.Exit(1)
	}

	return playerCountComputerAsInt, playerCountAsInt
}

func main() {

	globalGame = &game.Game{}
	globalGame.InitGame()

	// cheat settings
	globalGame.CheatCounter = 0
	globalGame.FirstCheatRound = 1
	globalGame.WinningPercent = 0.7
	globalGame.CheatsInARow = 0
	globalGame.NumberOfCheatMethods = 3

	computerPlayerCount, humanPlayerCount := getPlayerCount()

	for i := 0; i < humanPlayerCount; i++ {
		fmt.Println("==========================================================")
		globalGame.CreateHumanPlayer(i, "Enter player #"+strconv.Itoa(i+1)+" name: ..")
		fmt.Println("==========================================================")
		if globalGame.Players[i].Name == "AlphaDice" {
			globalGame.Players[i].IsCheater = true
		}
	}

	for i := humanPlayerCount; i < (computerPlayerCount + humanPlayerCount); i++ {
		globalGame.CreateComputerPlayer(i, "Computer-"+strconv.Itoa((i-humanPlayerCount)))
	}

	fmt.Println()
	fmt.Println("==========================================================")
	fmt.Println("Press r to roll the dice or press s to stop the game")
	fmt.Println("==========================================================")
	fmt.Println()
	startGame()
}

func startGame() {
NEXTROUND:
	// bump the game round
	globalGame.StartRound()
	globalGame.ResetAllDice()

	fmt.Println("==========================================================")
	fmt.Println("================= GAME ROUND: ", globalGame.Round)
	fmt.Println("==========================================================")

	// roll for all players
	for i := 0; i < len(globalGame.Players); i++ {
		fmt.Println("Next up:", globalGame.Players[globalGame.CurrentRoller].Name)
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
		fmt.Println(p.Name, "|", p.Wins, "| WinRating:", globalGame.WinRatings[i])
	}
	fmt.Println("Ties:", globalGame.TieCount)
}
