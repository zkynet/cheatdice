package main

import "testing"

func TestMain(t *testing.T) {

}

func TestDiceRoll1000000Times(t *testing.T) {
	min := 1
	max := 6
	for i := 0; i < 1000000; i++ {
		numb := rollDice(int64(min), int64(max))
		if int(numb) < min || int(numb) > max {
			t.Error("Expected a roll between ", min, "and", max, " but got :", numb)
			t.Fail()
		}
	}
}

func BenchmarkDiceRoll(b *testing.B) {
	min := 1
	max := 6
	for n := 0; n < b.N; n++ {
		_ = rollDice(int64(min), int64(max))
	}
}
