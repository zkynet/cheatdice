# Cheat Dice
A game of dice against the computer where the computer can be configured to sustain a certain win rate.

# Running the game
- go run main.go

# How the computer cheats
## When it cheats
The computer will start cheating once it passes the configured "FirstCheatRound". Once it has passed it will only cheat if it is under the set "WinningPercent".
## How it cheats
When it is the computers turn, it will roll a dice (dice inception intended) to determine one of three methods it can use to cheat. 
- The first method ( Dou    ble )
    - The computer performs a normal roll but will then set one dice to be equal to the other.
- The Second method ( Highest total )
    - the computer performs a normall roll but wil then check if any dice are lower then 3, if they are the set them to 3. Once both dice have been checked and possibly raised above 3, they get +1 to their value.
- The third method ( Highest Dice )
    - The computer performs a normall roll but then set's once dice to be equal to 6
## Why it seems random
1. This kind of cheating does not rely on reading the apponents dice, thus there is no 1-1 relation between the human dice roll and the cheat methods.
2. With multiple cheat methods it get's increasingly difficult to detect patterns
3. The cheating interval, win percentage and number of cheating method can be adjusted for each game. This makes it even more difficult to detect if optimized properly. 

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

# Roll log for 100 games 
```
33-46
36-11
44-41
62-16
13-11
64-11
33-23
44-26
33-64
11-42
61-41
24-61
66-54
63-32
54-24
66-26
63-45
62-63
64-34
31-61
65-66
65-55
62-42
45-25
46-16
11-23
65-56
65-22
62-54
64-22
61-32
64-23
44-54
11-53
55-15
55-31
63-31
46-64
44-51
66-65
46-16
64-26
55-31
65-13
56-61
32-51
66-62
63-22
44-66
63-13
45-51
66-12
55-62
22-24
32-62
66-16
66-44
32-46
61-25
61-31
65-31
35-21
14-51
33-66
66-65
65-52
33-53
45-56
66-31
65-55
22-52
61-35
63-61
64-24
45-56
65-41
11-21
61-16
22-43
64-55
62-32
33-44
63-13
62-46
66-51
46-13
64-11
64-55
22-62
61-44
64-15
64-23
56-53
44-51
55-24
64-11
55-45
44-61
66-55
62-64
```