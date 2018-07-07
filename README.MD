# Cheat Dice
A game of dice against the computer where the computer can be configured to sustain a certain win rate.

# Installing 
- go get github.com/zkynet/cheatdice

# Running the game
- go run main.go

# How the computer cheats
## When it cheats
The computer will start cheating once it passes the configured "FirstCheatRound". Once it has passed it will only cheat if it is under the set "WinningPercent".
## How it cheats
When it is the computers turn, it will roll a dice (dice inception intended) to determine one of three methods it can use to cheat. 
- The first method ( Double )
    - The computer performs a normal roll but will then set one dice to be equal to the other.
- The Second method ( Highest total )
    - the computer performs a normall roll but wil then check if any dice are lower then 3, if they are the set them to 3. Once both dice have been checked and possibly raised above 3, they get +1 to their value.
- The third method ( Highest Dice )
    - The computer performs a normall roll but then set's once dice to be equal to 5. I thought about using 6 here but the density of sixes in the roll log was too high with that setting. 5 seemed to fit better and still produce optimal win rates. 
## Why it seems random
1. This kind of cheating does not rely on reading the apponents dice, thus there is no 1-1 relation between the human dice roll and the cheat methods.
2. The cheating interval, win percentage and number of cheating method can be adjusted for each game. This makes it even more difficult to detect if optimized properly. 

# How to adjust cheat rate
```
type Game struct {
    ...
	CheatsInARow         int // how many times the computer can cheat in a row ( 0 means no limit )
	WinningPercent       float64 // the precentage of wins ( 0.1 to 0.9 )
	FirstCheatRound      int // when to start cheating
	NumberOfCheatMethods int // how any different methods to use for cheating
}
```

# Testing
- go test -v

# 1.000.000 game test with 70% win rate
settings
```
globalGame.FirstCheatRound = 1
globalGame.WinningPercent = 0.7
globalGame.CheatsInARow = 0
globalGame.NumberOfCheatMethods = 3
```
output
```
=== RUN   TestWinningPercentage1000000Rounds
Win rating: 0.699997
--- PASS: TestWinningPercentage1000000Rounds (10.14s)
```
# 1.000.000 game test with 80% win rate
settings
```
globalGame.FirstCheatRound = 1
globalGame.WinningPercent = 0.8
globalGame.CheatsInARow = 0
globalGame.NumberOfCheatMethods = 3
```
output
```
=== RUN   TestWinningPercentage1000000Rounds
Win rating: 0.7957014298570143
--- PASS: TestWinningPercentage1000000Rounds (10.37s)
```
# 1.000.000 game test with 90% win rate
90% win rating is hard to reach while making it seem random. To get to 90% I lower the number of methods to be used for cheating. By default the weakest method is removed first.

 - I only get up to about 85% win rate with 2 methods.

settings
```
globalGame.FirstCheatRound = 1
globalGame.WinningPercent = 0.9
globalGame.CheatsInARow = 0
globalGame.NumberOfCheatMethods = 2
```
output
```
=== RUN   TestWinningPercentage1000000Rounds
Win rating: 0.8492200779922008
--- PASS: TestWinningPercentage1000000Rounds (12.38s)
```

- Using only the best win method we can reach 90%, but it's not looking very random at this point.

settings
```
globalGame.FirstCheatRound = 1
globalGame.WinningPercent = 0.9
globalGame.CheatsInARow = 0
globalGame.NumberOfCheatMethods = 1
```
output
```
=== RUN   TestWinningPercentage1000000Rounds
Win rating: 0.89999600039996
--- PASS: TestWinningPercentage1000000Rounds (11.35s)
```

# 100 game test
settings
```
globalGame.FirstCheatRound = 1
globalGame.WinningPercent = 0.7
globalGame.CheatsInARow = 0
globalGame.NumberOfCheatMethods = 3
```
output
```
=== RUN   TestWinningPercentage100Rounds
Win rating: 0.67
--- PASS: TestWinningPercentage100Rounds (0.00s)
```

# Example test suit output
```
zkynets-MacBook-Pro:cheatdice zkynet$ go test -v --bench=.
=== RUN   TestDiceRoll3000000TimesMax6
--- PASS: TestDiceRoll3000000TimesMax6 (2.96s)
=== RUN   TestDiceRollWithATie
=== RUN   TestDiceRollWithATie/Players_tie
=== RUN   TestDiceRollWithATie/Players_tie#01
--- PASS: TestDiceRollWithATie (0.00s)
    --- PASS: TestDiceRollWithATie/Players_tie (0.00s)
    --- PASS: TestDiceRollWithATie/Players_tie#01 (0.00s)
=== RUN   TestDiceRollWinMethodHighestDice
=== RUN   TestDiceRollWinMethodHighestDice/Player_0_wins_with_the_highest_dice
=== RUN   TestDiceRollWinMethodHighestDice/Player_0_wins_with_the_highest_dice#01
=== RUN   TestDiceRollWinMethodHighestDice/Player_0_wins_with_the_highest_dice#02
=== RUN   TestDiceRollWinMethodHighestDice/Player_1_wins_with_the_highest_dice
--- PASS: TestDiceRollWinMethodHighestDice (0.00s)
    --- PASS: TestDiceRollWinMethodHighestDice/Player_0_wins_with_the_highest_dice (0.00s)
    --- PASS: TestDiceRollWinMethodHighestDice/Player_0_wins_with_the_highest_dice#01 (0.00s)
    --- PASS: TestDiceRollWinMethodHighestDice/Player_0_wins_with_the_highest_dice#02 (0.00s)
    --- PASS: TestDiceRollWinMethodHighestDice/Player_1_wins_with_the_highest_dice (0.00s)
=== RUN   TestDiceRollWinMethodHigherTotal
=== RUN   TestDiceRollWinMethodHigherTotal/Player_0_wins_with_a_higher_total
=== RUN   TestDiceRollWinMethodHigherTotal/Player_1_wins_with_a_higher_total
--- PASS: TestDiceRollWinMethodHigherTotal (0.00s)
    --- PASS: TestDiceRollWinMethodHigherTotal/Player_0_wins_with_a_higher_total (0.00s)
    --- PASS: TestDiceRollWinMethodHigherTotal/Player_1_wins_with_a_higher_total (0.00s)
=== RUN   TestDiceRollWinMethodDouble
=== RUN   TestDiceRollWinMethodDouble/Player_0_wins_with_a_double
=== RUN   TestDiceRollWinMethodDouble/Player_1_wins_with_a_double
=== RUN   TestDiceRollWinMethodDouble/Player_0_wins_with_a_higher_double
=== RUN   TestDiceRollWinMethodDouble/Player_1_wins_with_a_higher_double
--- PASS: TestDiceRollWinMethodDouble (0.00s)
    --- PASS: TestDiceRollWinMethodDouble/Player_0_wins_with_a_double (0.00s)
    --- PASS: TestDiceRollWinMethodDouble/Player_1_wins_with_a_double (0.00s)
    --- PASS: TestDiceRollWinMethodDouble/Player_0_wins_with_a_higher_double (0.00s)
    --- PASS: TestDiceRollWinMethodDouble/Player_1_wins_with_a_higher_double (0.00s)
=== RUN   TestWinningPercentage100Rounds70WinRatingWithIn10Percent
--- PASS: TestWinningPercentage100Rounds70WinRatingWithIn10Percent (0.00s)
=== RUN   TestWinningPercentage1000000Rounds70WinRatingWithIn10Percent
--- PASS: TestWinningPercentage1000000Rounds70WinRatingWithIn10Percent (10.95s)
=== RUN   TestWinningPercentage1000000Rounds70WinRatingWithIn5Percent
--- PASS: TestWinningPercentage1000000Rounds70WinRatingWithIn5Percent (3.38s)
goos: darwin
goarch: amd64
pkg: github.com/zkynet/cheatdice
BenchmarkDiceRoll-4      2000000               960 ns/op
PASS
ok      github.com/zkynet/cheatdice     20.172s
zkynets-MacBook-Pro:cheatdice zkynet$
```

# Roll log for 100 games 
```
44-34
33-24
24-44
11-31
31-45
22-45
56-33
44-64
61-23
52-55
56-14
42-21
44-11
45-26
66-13
64-22
62-51
25-24
45-54
54-36
53-54
46-33
45-65
56-41
34-64
64-24
66-46
33-11
52-16
66-24
45-24
45-46
64-61
45-16
64-25
66-21
34-22
46-22
56-42
52-46
44-42
33-14
23-65
55-34
54-25
45-34
44-53
64-53
62-41
22-25
66-12
14-56
64-46
53-44
46-14
44-22
45-32
35-11
44-56
46-43
63-63
52-45
64-11
56-64
56-24
64-52
62-11
51-45
56-51
14-12
66-31
56-45
42-32
56-52
11-61
55-61
66-22
45-61
65-54
56-43
22-65
22-22
51-25
44-64
25-41
44-33
56-22
26-31
52-31
66-55
43-41
64-13
46-65
34-52
55-16
22-36
12-41
54-46
54-42
61-51
```