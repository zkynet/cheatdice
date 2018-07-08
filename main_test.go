package main

import (
	"testing"

	game "github.com/zkynet/cheatdice/Game"
)

func newGame() {
	computer0 := game.Player{
		Name:        "AlphaDice0",
		Wins:        0,
		Rolls:       0,
		IsComputer:  true,
		IsCheater:   true,
		CurrentDice: make(map[int]int),
	}

	computer1 := game.Player{
		Name:        "AlphaDice1",
		Wins:        0,
		Rolls:       0,
		IsComputer:  true,
		IsCheater:   false,
		CurrentDice: make(map[int]int),
	}
	globalGame = &game.Game{}
	globalGame.InitGame()
	globalGame.Players[0] = &computer0
	globalGame.Players[1] = &computer1
	globalGame.CheatCounter = 0
	globalGame.FirstCheatRound = 1
	globalGame.WinningPercent = 0.7
	globalGame.CheatsInARow = 0
	globalGame.NumberOfCheatMethods = 3
}

func TestDiceRoll3000000TimesMax6(t *testing.T) {
	p := game.Player{}
	p.CurrentDice = make(map[int]int)
	max := 6
	for i := 0; i < 3000000; i++ {
		p.ResetDice()
		p.RollDice(max)
		if p.CurrentDice[0] > max {
			t.Error("Expected a roll lower then", max, " but got :", p.CurrentDice[0])
			t.Fail()
		}
		if p.CurrentDice[1] > max {
			t.Error("Expected a roll lower then", max, " but got :", p.CurrentDice[1])
			t.Fail()
		}
	}
}

func testDiceOutcome(t *testing.T, winner *game.Player, message string, expectedWinnerName string, expectedMessage string) {
	if winner == nil {
		t.Error("Expected a winner but got a tie")
		t.Fail()
		return
	}

	if winner.Name != expectedWinnerName {
		t.Error("Expected winner", expectedWinnerName, " --- but got", winner.Name)
		t.Fail()
	}

	if message != expectedMessage {
		t.Error("Expected message:", expectedMessage, " --- but got:", message)
		t.Fail()
	}
}

func TestDiceRollWithATie(t *testing.T) {
	newGame()

	t.Run("Players tie", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 2
		globalGame.Players[0].CurrentDice[1] = 3
		globalGame.Players[1].CurrentDice[0] = 2
		globalGame.Players[1].CurrentDice[1] = 3

		winner, message := globalGame.FindRoundWinner()
		if winner != nil {
			t.Error("Excpected a tie but got message:", message)
		}
	})

	t.Run("Players tie", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 4
		globalGame.Players[0].CurrentDice[1] = 4
		globalGame.Players[1].CurrentDice[0] = 4
		globalGame.Players[1].CurrentDice[1] = 4

		winner, message := globalGame.FindRoundWinner()
		if winner != nil {
			t.Error("Excpected a tie but got message:", message)
		}
	})

}

func TestDiceRollWinMethodHighestDice(t *testing.T) {
	newGame()

	t.Run("Player 0 wins with the highest dice", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 4
		globalGame.Players[1].CurrentDice[0] = 3
		globalGame.Players[1].CurrentDice[1] = 2

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[0].Name, "Player "+globalGame.Players[0].Name+" won with the highest dice")
	})

	t.Run("Player 0 wins with the highest dice", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 5
		globalGame.Players[1].CurrentDice[0] = 4
		globalGame.Players[1].CurrentDice[1] = 2

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[0].Name, "Player "+globalGame.Players[0].Name+" won with the highest dice")
	})

	t.Run("Player 0 wins with the highest dice", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 6
		globalGame.Players[1].CurrentDice[0] = 5
		globalGame.Players[1].CurrentDice[1] = 2

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[0].Name, "Player "+globalGame.Players[0].Name+" won with the highest dice")
	})

	t.Run("Player 1 wins with the highest dice", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 2
		globalGame.Players[0].CurrentDice[1] = 3
		globalGame.Players[1].CurrentDice[0] = 1
		globalGame.Players[1].CurrentDice[1] = 4

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[1].Name, "Player "+globalGame.Players[1].Name+" won with the highest dice")
	})
}
func TestDiceRollWinMethodHigherTotal(t *testing.T) {
	newGame()

	t.Run("Player 0 wins with a higher total", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 3
		globalGame.Players[1].CurrentDice[0] = 1
		globalGame.Players[1].CurrentDice[1] = 2

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[0].Name, "Player "+globalGame.Players[0].Name+" won with a higher total")
	})

	t.Run("Player 1 wins with a higher total", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 2
		globalGame.Players[1].CurrentDice[0] = 1
		globalGame.Players[1].CurrentDice[1] = 3

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[1].Name, "Player "+globalGame.Players[1].Name+" won with a higher total")
	})
}
func TestDiceRollWinMethodDouble(t *testing.T) {
	newGame()

	t.Run("Player 0 wins with a double", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 1
		globalGame.Players[1].CurrentDice[0] = 1
		globalGame.Players[1].CurrentDice[1] = 2

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[0].Name, "Player "+globalGame.Players[0].Name+" won with the biggest double")
	})

	t.Run("Player 1 wins with a double", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 2
		globalGame.Players[1].CurrentDice[0] = 1
		globalGame.Players[1].CurrentDice[1] = 1

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[1].Name, "Player "+globalGame.Players[1].Name+" won with the biggest double")
	})

	t.Run("Player 0 wins with a higher double", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 2
		globalGame.Players[0].CurrentDice[1] = 2
		globalGame.Players[1].CurrentDice[0] = 1
		globalGame.Players[1].CurrentDice[1] = 1

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[0].Name, "Player "+globalGame.Players[0].Name+" won with the biggest double")
	})

	t.Run("Player 1 wins with a higher double", func(t *testing.T) {
		globalGame.Players[0].CurrentDice[0] = 1
		globalGame.Players[0].CurrentDice[1] = 1
		globalGame.Players[1].CurrentDice[0] = 2
		globalGame.Players[1].CurrentDice[1] = 2

		winner, message := globalGame.FindRoundWinner()
		testDiceOutcome(t, winner, message, globalGame.Players[1].Name, "Player "+globalGame.Players[1].Name+" won with the biggest double")
	})

}

func TestWinningPercentage100Rounds70WinRatingWithIn10Percent(t *testing.T) {
	newGame()
	//f, err := os.OpenFile("100-round-roll-log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	//if err != nil {
	//	panic(err)
	//}

	//defer f.Close()

	for i := 0; i < 100; i++ {
		globalGame.Round = globalGame.Round + 1
		globalGame.ResetAllDice()
		globalGame.AskPlayerToRoll()
		globalGame.SwitchPlayers()
		globalGame.AskPlayerToRoll()
		_, _ = globalGame.FindRoundWinner()
		globalGame.CalculateWinRatings()

		//text := strconv.Itoa(globalGame.Players[0].CurrentDice[0]) + strconv.Itoa(globalGame.Players[0].CurrentDice[1]) + "-" + strconv.Itoa(globalGame.Players[1].CurrentDice[0]) + strconv.Itoa(globalGame.Players[1].CurrentDice[1]) + "\n"
		//if _, err = f.WriteString(text); err != nil {
		//	panic(err)
		//}
	}

	if globalGame.WinRatings[0] < 0.60 || globalGame.WinRatings[0] > 0.80 {
		t.Error("Win rating was:", globalGame.WinRatings[0], "wanted a value between 0.6 and 0.8")
		t.Fail()
	}

	//fmt.Println("Win rating:", globalGame.WinRatings[0])

}

func TestWinningPercentage1000000Rounds70WinRatingWithIn10Percent(t *testing.T) {
	newGame()
	//f, err := os.OpenFile("1000000-round-roll-log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	//if err != nil {
	//	panic(err)
	//}

	//defer f.Close()

	for i := 0; i < 1000000; i++ {
		globalGame.Round = globalGame.Round + 1
		globalGame.ResetAllDice()
		globalGame.AskPlayerToRoll()
		globalGame.SwitchPlayers()
		globalGame.AskPlayerToRoll()
		_, _ = globalGame.FindRoundWinner()
		globalGame.CalculateWinRatings()

		//text := strconv.Itoa(globalGame.Players[0].CurrentDice[0]) + strconv.Itoa(globalGame.Players[0].CurrentDice[1]) + "-" + strconv.Itoa(globalGame.Players[1].CurrentDice[0]) + strconv.Itoa(globalGame.Players[1].CurrentDice[1]) + "\n"
		//if _, err = f.WriteString(text); err != nil {
		//	panic(err)
		//}
	}

	if globalGame.WinRatings[0] < 0.60 || globalGame.WinRatings[0] > 0.80 {
		t.Error("Win rating was:", globalGame.WinRatings[0], "wanted a value between 0.6 and 0.8")
		t.Fail()
	}

	//fmt.Println("Win rating:", globalGame.WinRatings[0])

}

func TestWinningPercentage1000000Rounds70WinRatingWithIn5Percent(t *testing.T) {
	newGame()
	for i := 0; i < 1000000; i++ {
		globalGame.Round = globalGame.Round + 1
		globalGame.ResetAllDice()
		globalGame.AskPlayerToRoll()
		globalGame.SwitchPlayers()
		globalGame.AskPlayerToRoll()
		_, _ = globalGame.FindRoundWinner()
		globalGame.CalculateWinRatings()

	}

	if globalGame.WinRatings[0] < 0.65 || globalGame.WinRatings[0] > 0.75 {
		t.Error("Win rating was:", globalGame.WinRatings[0], "wanted a value between 0.65 and 0.75")
		t.Fail()
	}

	//fmt.Println("Win rating:", globalGame.WinRatings[0])

}

func BenchmarkDiceRoll(b *testing.B) {

	max := 6
	p := game.Player{}
	p.CurrentDice = make(map[int]int)

	for n := 0; n < b.N; n++ {
		p.RollDice(max)
	}
}
