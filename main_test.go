package main

import (
	"reflect"
	"testing"
)

func TestCalculatePlayersMoney(t *testing.T) {
	numPlayers := 3
	betMoney := 100
	wins := []int{2, 1, 3}
	expected := []Player{
		{ID: 1, Money: 100},
		{ID: 2, Money: -200},
		{ID: 3, Money: 100},
	}

	players := calculatePlayersMoney(numPlayers, betMoney, wins)
	if !reflect.DeepEqual(players, expected) {
		t.Errorf("calculatePlayersMoney() = %v; want %v", players, expected)
	}
}

func TestCalculateDebts(t *testing.T) {
	players := []Player{
		{ID: 1, Money: -200},
		{ID: 2, Money: 300},
		{ID: 3, Money: -100},
	}
	expected := []string{
		"Player 1 pays Player 2 $200",
		"Player 3 pays Player 2 $100",
	}

	payments := calculateDebts(players)
	if !reflect.DeepEqual(payments, expected) {
		t.Errorf("calculateDebts() = %v; want %v", payments, expected)
	}
}
