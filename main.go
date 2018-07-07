package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

type Player struct {
	Number         int
	Name           string
	Wins           int
	Rolls          int
	isComputer     bool // if true, rolls will happen automaticaly
	isCheater      bool // if true, this player will get software assistance
	CheatCounter   int
	WinningPercent int
	Dice           map[int]int
}

func (p *Player) RollDice(min int, max int) {
	p.Dice[0] = int(rollDice(int64(min), int64(max)))
	p.Dice[1] = int(rollDice(int64(min), int64(max)))
}

func (p *Player) ResetDice() {
	p.Dice[0] = 0
	p.Dice[1] = 0
}

type Game struct {
	Players    map[int]*Player
	NextToRoll int
	Round      int
	Dice       *Dice
}

type Dice struct {
	Min int
	Max int
}

func (g *Game) Cheat() bool {
	return true
}
func (g *Game) AnnounceRoundWinner() int {
	// rules
	// 1. double or higher double wins.
	// Set some variables for readability
	player0Dice0 := game.Players[0].Dice[0]
	player0Dice1 := game.Players[0].Dice[1]
	player1Dice0 := game.Players[1].Dice[0]
	player1Dice1 := game.Players[1].Dice[1]

	// if player 0 has a double and player 1 does not
	if player0Dice0 == player0Dice1 && player1Dice0 != player1Dice1 {
		return game.Players[0].Number
	}

	// if player 1 has a double and player 0 does not
	if player1Dice0 == player1Dice1 && player0Dice0 != player0Dice1 {
		return game.Players[1].Number
	}

	// if both players have doubles
	if game.Players[1].Dice[0] == game.Players[1].Dice[1] &&
		game.Players[0].Dice[0] == game.Players[0].Dice[1] {
			if game.Players[1].Dice[0]
		return game.Players[1].Number
	}

	// 2. higher total wins
	// 3. highest number wins if total is the same
	// 4. draw, no points.

	fmt.Println(winner)
}

// global game struct .. etc ..
var game *Game

func main() {
	game = initGame()
	// Game loop
	// 1. read input
	// 1.1 if it is the computers turn we need to analyze score and determine if we need to cheat
	// -- if we cheat we can only do it when the player goes first
	// -- needs to win X percent of games
	// -- can never cheat three in a row
	// 2. roll or stop
	// 3.

	fmt.Println(game)
	fmt.Println("Press r to roll the dice or press s to stop the game")

LOOP:
	game.Round = game.Round + 1
	fmt.Println("======== GAME ROUND: ", game.Round, "=========")
	nextRound()
	printStandings()
	fmt.Println("========= END OF ROUND =========")
	goto LOOP
}

func initGame() *Game {

	fmt.Println("Human player name:")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Could not parse your name, err:", err)
		os.Exit(1)
	}

	computer := Player{
		Number:         0,
		Name:           "AlphaDice:)",
		Wins:           0,
		Rolls:          0,
		isComputer:     true,
		isCheater:      true,
		WinningPercent: 70,
		CheatCounter:   0,
		Dice:           make(map[int]int),
	}

	human := Player{
		Number:     1,
		Name:       string(name),
		Wins:       0,
		Rolls:      0,
		isComputer: false,
		isCheater:  false,
		Dice:       make(map[int]int),
	}

	game := Game{
		NextToRoll: 0,
		Round:      0,
		Players:    make(map[int]*Player),
		Dice: &Dice{
			Min: 1,
			Max: 6,
		},
	}
	game.Players[0] = &computer
	game.Players[1] = &human

	return &game
}

func nextRound() {
	rollDiceForPlayer(game.NextToRoll)
	printDiceRollForPlayer(game.NextToRoll)
	switchPlayers()
	rollDiceForPlayer(game.NextToRoll)
	printDiceRollForPlayer(game.NextToRoll)
	// Check whom wins.

}

func userInputPrompt() {
	// do not move outside of loop, this will cause new line(enter) to be read as the next character
	reader := bufio.NewReader(os.Stdin)

PROMPTLOOP:
	switch readChar(reader) {
	case 'R', 'r':
		return
	case 'S', 's':
		fmt.Println()
		printResults()
		fmt.Println()
		os.Exit(1)
	default:
		fmt.Println("You pressed something other then r or s, please try again..")
		goto PROMPTLOOP
	}

}

func rollDiceForPlayer(number int) {

	if game.Players[number].isComputer {
		if game.Cheat() {
			// if the player is a computer and it's cheating then we don't need to roll the dice.
			return
		}
	} else {
		userInputPrompt()
	}

	game.Players[number].ResetDice()
	game.Players[number].RollDice(game.Dice.Min, game.Dice.Max)
}

func switchPlayers() {
	if game.NextToRoll == 0 {
		game.NextToRoll = 1
	} else {
		game.NextToRoll = 0
	}
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

func rollDice(min int64, max int64) int64 {
	result, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		fmt.Println("Can not generate random numbers.. game halted. Error: ", err)
		os.Exit(1)
	}
	return result.Int64() + min
}

// ====================================
//
// COSMETIC
//
// ====================================

func printDiceRollForPlayer(number int) {
	fmt.Println("Dice Roll: ", game.Players[number].Name, " (", game.Players[number].Dice[0], "/", game.Players[number].Dice[1], ")")
}

func printStandings() {
	fmt.Println("game standings")
}

func printResults() {
	fmt.Println("game results")

}
