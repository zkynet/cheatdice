package game

import (
	"bufio"
	"fmt"
	"os"
)

type Game struct {
	Players              map[int]*Player // player list
	CurrentRoller        int             // the current dice roller
	StarrtingRoller      int             // The player that starts
	Round                int             // The current round
	Dice                 *Dice           // The current dice settings for this game
	WinRatings           map[int]float64 // The current win ratings of the game
	TieCount             int             // How many ties have happened
	CheatCounter         int             // how many cheats have been used in a row
	CheatsInARow         int             // how many cheats can be used in a row
	WinningPercent       float64         // how much percentage should the cheater win by ? (0.1-0.9)
	FirstCheatRound      int             // When to start cheating
	NumberOfCheatMethods int             // how many methods should be used to cheat ?
}

func (g *Game) AskPlayerToRoll() {

	if !g.Players[g.CurrentRoller].IsComputer {
		// if the player is not a computer we check for input
		userInputPrompt()
	}

	if g.Cheat() {
		// if the player is a cheater we do not roll normally
		return
	}

	g.Players[g.CurrentRoller].ResetDice()
	g.Players[g.CurrentRoller].RollDice(g.Dice.Max)
}

func (g *Game) Cheat() bool {
	// only start cheating after ten rounds
	if g.Round < g.FirstCheatRound {
		return false
	}

	currentRoller := g.Players[g.CurrentRoller]
	// do not do anything if the player is not a cheater
	if !currentRoller.IsCheater {
		return false
	}

	// if we are above the winning percentile we do not cheat
	if g.WinRatings[g.CurrentRoller] > g.WinningPercent {
		return false
	}

	// only cheat a certain times in a row
	// always cheat if 0
	if g.CheatCounter == g.CheatsInARow && g.CheatsInARow != 0 {
		g.CheatCounter = 0
		return false
	}

	// roll a dice to determine which method we will cheat with
	g.Players[g.CurrentRoller].RollDice(g.NumberOfCheatMethods)
	switch currentRoller.CurrentDice[0] {

	case 1: // win with a double
		currentRoller.RollDice(6)
		// if dice 1 is higher then 4 lower it by 1
		if currentRoller.CurrentDice[1] > 4 {
			currentRoller.CurrentDice[1]--
		}
		// set dice 0 to be the same as 1
		currentRoller.CurrentDice[0] = currentRoller.CurrentDice[1]
		break

	case 2: // win with a higher total
		currentRoller.RollDice(6)
		// if the current dice are below 3, bump them up to 3
		if currentRoller.CurrentDice[0] < 3 {
			currentRoller.CurrentDice[0] = 3
		}
		if currentRoller.CurrentDice[1] < 3 {
			currentRoller.CurrentDice[1] = 3
		}
		// increase both dice by 1
		if currentRoller.CurrentDice[0] < 6 {
			currentRoller.CurrentDice[0]++
		}
		if currentRoller.CurrentDice[1] < 6 {
			currentRoller.CurrentDice[1]++
		}
		break

	case 3: // win with highest dice
		currentRoller.RollDice(6)
		// if dice 0 is higher then 3 set it to 5, otherwise set dice 1 to 5
		if currentRoller.CurrentDice[0] > 3 {
			currentRoller.CurrentDice[0] = 5
		} else {
			currentRoller.CurrentDice[1] = 5
		}
		break
	}

	// do not increase counter if we are not cheating
	if g.CheatsInARow != 0 {
		g.CheatCounter++
	}

	return true
}

func (g *Game) ResetAllDice() {
	for _, player := range g.Players {
		player.ResetDice()
	}
}

func (g *Game) SwitchPlayers() {
	if g.CurrentRoller+1 == len(g.Players) {
		g.CurrentRoller = 0
	} else {
		g.CurrentRoller++
	}
}

func (g *Game) SwitchStartingPlayer() {
	if g.StarrtingRoller+1 == len(g.Players) {
		g.CurrentRoller = 0
		g.StarrtingRoller = 0
	} else {
		g.StarrtingRoller++
		g.CurrentRoller = g.StarrtingRoller
	}
}

func (g *Game) StartRound() {
	g.Round++
}

func (g *Game) hasEveryoneRolled() bool {
	for _, p := range g.Players {
		if p.CurrentDice[0] == 0 || p.CurrentDice[1] == 0 {
			return false
		}
	}
	return true
}

func (g *Game) findHighestTotal() (highestTotalOwnerIndex int) {
	highestTotal := 0
	highestTotalCount := 0

	// for all the players in the game
	for playerIndex, player := range g.Players {
		// get the current players total
		total := player.CurrentDice[0] + player.CurrentDice[1]
		// if the total is higher then the highest total
		if total > highestTotal {
			highestTotal = total                 // set it to highest total
			highestTotalOwnerIndex = playerIndex // save the current players index
			highestTotalCount = 1                // set the highest count to one
		}

		// if the current players total is the same as the highest
		if total == highestTotal &&
			// and they do not have the same owner
			playerIndex != highestTotalOwnerIndex {
			// increase the highest total counter by one
			highestTotalCount++
		}

	}

	if highestTotalCount == 1 {
		// if we only have one player with the highest total we return
		return
	}

	// if there were more players then one with the highest total
	// we set the player index to -1
	highestTotalOwnerIndex = -1
	return
}

func (g *Game) findHighestDice() (highestDiceOwnerIndex int) {
	highestDiceValue := 0
	highestDiceCount := 0

	// for all the players in the game
	for playerIndex, player := range g.Players {
		// and all of their dice
		for _, diceValue := range player.CurrentDice {
			// if the current dice is higher then the highest dice value
			if diceValue > highestDiceValue {
				highestDiceValue = diceValue        // set the current dice as the highest
				highestDiceOwnerIndex = playerIndex // save the player index
				highestDiceCount = 1                // set the highest count to one
			}

			// if the current dice is the same as the current highest dice
			if diceValue == highestDiceValue &&
				// and they do not have the same owner
				highestDiceOwnerIndex != playerIndex {
				// increase the highest dice counter by one
				highestDiceCount++
			}
		}
	}

	if highestDiceCount == 1 {
		// if we only have one player with the highest dice we return
		return
	}

	// if there were more players then one with the highest dice
	// we set the player index to -1
	highestDiceOwnerIndex = -1
	return

}

func (g *Game) findHighestDouble() (highestDoubleOwnerIndex int) {
	highestDouble := 0
	highestDoubleCount := 0
	// for all the player in the game
	for playerIndex, player := range g.Players {
		// check if the current dice pair is the same
		if player.CurrentDice[0] == player.CurrentDice[1] {
			//if the current dice is higher then the highest double dice
			if player.CurrentDice[0] > highestDouble {
				highestDouble = player.CurrentDice[0] // set it as the highest dice
				highestDoubleOwnerIndex = playerIndex // save the current player index
				highestDoubleCount = 1                // set the highest count to one
			}

			// if the current players dice is the same as the highest
			if player.CurrentDice[0] == highestDouble &&
				// and they do not have the same owner
				playerIndex != highestDoubleOwnerIndex {
				// increase the highest duble counter by one
				highestDoubleCount++
			}

		}
	}

	if highestDoubleCount == 1 {
		// if we only have one player with the highest double we return
		return
	}

	// if there were more players then one with the highest double
	// we set the player index to -1
	highestDoubleOwnerIndex = -1
	return
}

func (g *Game) FindRoundWinner() (winner *Player, message string) {

	currentWinnerIndex := g.findHighestDouble()
	if currentWinnerIndex > -1 {
		winner = g.Players[currentWinnerIndex]
		message = "Player " + winner.Name + " won with the biggest double"
		winner.Wins++
		return
	}

	currentWinnerIndex = g.findHighestTotal()
	if currentWinnerIndex > -1 {
		winner = g.Players[currentWinnerIndex]
		message = "Player " + winner.Name + " won with a higher total"
		winner.Wins++
		return
	}

	currentWinnerIndex = g.findHighestDice()
	if currentWinnerIndex > -1 {
		winner = g.Players[currentWinnerIndex]
		message = "Player " + winner.Name + " won with the highest dice"
		winner.Wins++
		return
	}

	// The game ended in a draw!
	g.TieCount++
	winner = nil
	message = "The game was a tie"
	return
}

func (g *Game) InitGame() {

	g.CurrentRoller = 0
	g.StarrtingRoller = 0
	g.Round = 0
	g.Players = make(map[int]*Player)
	g.WinRatings = make(map[int]float64)
	g.Dice = &Dice{
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

func readChar(reader *bufio.Reader) rune {
READ:
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println("An unexcpected error occurred: ", err)
		goto READ
	}

	return char
}
