package game

import (
	"fmt"
)

type Game struct {
	Players        map[int]*Player
	NextToRoll     int
	Round          int
	Dice           *Dice
	WinRatings     map[int]float64
	CheatCounter   int
	WinningPercent float64
	CheatOnRound   int
}

func (g *Game) Cheat() bool {
	// only start cheating after ten rounds
	if g.Round < g.CheatOnRound {
		return false
	}

	// do not do anything if the player is not a cheater
	if !g.Players[g.NextToRoll].IsCheater {
		return false
	}

	if g.WinRatings[g.NextToRoll] < g.WinningPercent {
		fmt.Println("we should cheat now !")
	}

	return true
}

func (g *Game) ResetAllDice() {
	for _, p := range g.Players {
		p.ResetDice()
	}
}

func (g *Game) SwitchPlayers() {
	if g.NextToRoll == 0 {
		g.NextToRoll = 1
	} else {
		g.NextToRoll = 0
	}

}

func (g *Game) FindRoundWinner() (winner *Player, message string) {

	// Set some variables for readability
	player0Dice0 := g.Players[0].DiceRolls[0]
	player0Dice1 := g.Players[0].DiceRolls[1]
	player1Dice0 := g.Players[1].DiceRolls[0]
	player1Dice1 := g.Players[1].DiceRolls[1]
	player1total := player1Dice0 + player1Dice1
	player0total := player0Dice0 + player0Dice1

	// if player 0 has a double and player 1 does not
	if player0Dice0 == player0Dice1 && player1Dice0 != player1Dice1 {
		winner = g.Players[0]
		winner.Wins++
		message = "Player " + winner.Name + " won with a double"
		return
	}

	// if player 1 has a double and player 0 does not
	if player1Dice0 == player1Dice1 && player0Dice0 != player0Dice1 {
		winner = g.Players[1]
		winner.Wins++
		message = "Player " + winner.Name + " won with a double"
		return
	}

	// if both players have doubles
	if player1Dice0 == player1Dice1 && player0Dice0 == player0Dice1 {
		if player0Dice0 > player1Dice0 {
			// if player0 has a bigger double
			winner = g.Players[0]
			winner.Wins++
			message = "Player " + winner.Name + " won with a higher double"
			return
		}
		// if player1 has a bigger double
		winner = g.Players[1]
		winner.Wins++
		message = "Player " + winner.Name + " won with a higher double"
		return
	}

	// player 1 has a higher total
	if player1total > player0total {
		winner = g.Players[1]
		winner.Wins++
		message = "Player " + winner.Name + " won with a higher total"
		return
	}

	// Player 0 has a higher total
	if player1total < player0total {
		winner = g.Players[0]
		winner.Wins++
		message = "Player " + winner.Name + " won with a higher total"
		return
	}

	// If both players have the same totals
	if player1total == player0total {
		highestPlayer := g.findHighestDice()
		// -1 means a tie, anything higher then that represents a player index number
		if highestPlayer > -1 {
			winner = g.Players[highestPlayer]
			winner.Wins++
			message = "Player " + winner.Name + " won with the highest dice"
			return
		}
	}

	// The game ended in a draw!
	winner = nil
	message = "The game was a tie"
	return
}

func (g *Game) findHighestDice() int {
	var highestDiceValue = 0
	var highestDiceCount = 0
	var highestDiceOwnerIndex = -1
	// check all the dice from all the players
	for playerIndex, player := range g.Players {
		for _, diceValue := range player.DiceRolls {
			// if the dice value is higher then the current highest value
			if diceValue > highestDiceValue {
				// and the dice does not belong to the current highest value owner
				if highestDiceOwnerIndex != playerIndex {
					highestDiceValue = diceValue
					highestDiceOwnerIndex = playerIndex
					highestDiceCount = 1
				}
			}

			// if the dice values are the same
			if diceValue == highestDiceValue &&
				// but there are different owners
				highestDiceOwnerIndex != playerIndex {
				highestDiceCount++
			}
		}
	}

	if highestDiceCount > 1 {
		// if we have more then one player with the same highest dice count we have a tie
		return -1
	}

	return highestDiceOwnerIndex

}

func (g *Game) InitGame() {

	g.NextToRoll = 0
	g.Round = 0
	g.Players = make(map[int]*Player)
	g.WinRatings = make(map[int]float64)
	g.Dice = &Dice{
		Min: 1,
		Max: 6,
	}

}

func (g *Game) CalculateWinRatings() {
	for i, p := range g.Players {
		if p.Wins > 0 {
			g.WinRatings[i] = float64(p.Wins) / float64(g.Round)
		}
	}
}
