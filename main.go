package main

import (
	"fmt"
	"sort"
)

type Player struct {
	ID    int
	Money int
}

func main() {
	numPlayers, betMoney, wins := inputPlayersData()
	players := calculatePlayersMoney(numPlayers, betMoney, wins)
	displayPlayersMoney(players)
	payments := calculateDebts(players)
	displayPayments(payments)
}

func inputPlayersData() (int, int, []int) {
	var numPlayers, betMoney int

	fmt.Print("Enter the number of players: ")
	fmt.Scanln(&numPlayers)
	fmt.Print("Enter the betting money for each round: ")
	fmt.Scanln(&betMoney)

	wins := make([]int, numPlayers)
	totalRounds := 0
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter the total winning rounds for player %d: ", i+1)
		fmt.Scanln(&wins[i])
		totalRounds += wins[i]
	}
	return numPlayers, betMoney, wins
}

func calculatePlayersMoney(numPlayers, betMoney int, wins []int) []Player {
	totalRounds := 0
	for _, win := range wins {
		totalRounds += win
	}

	players := make([]Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = Player{
			ID:    i + 1,
			Money: (wins[i] * (numPlayers - 1) * betMoney) - ((totalRounds - wins[i]) * betMoney),
		}
	}
	return players
}

func displayPlayersMoney(players []Player) {
	fmt.Println("Final money for each player:")
	for _, player := range players {
		fmt.Printf("Player %d: %d\n", player.ID, player.Money)
	}
}

func calculateDebts(players []Player) []string {
	sort.Slice(players, func(i, j int) bool {
		return players[i].Money < players[j].Money
	})

	var payments []string
	for i, j := 0, len(players)-1; i < j; {
		playerA, playerB := &players[i], &players[j]
		if playerA.Money >= 0 {
			break
		}

		transfer := playerA.Money
		if playerB.Money+transfer >= 0 {
			playerA.Money += transfer
			playerB.Money += transfer
			payments = append(payments, fmt.Sprintf("Player %d pays Player %d $%d", playerA.ID, playerB.ID, -transfer))
			i++
		} else {
			playerA.Money += playerB.Money
			payments = append(payments, fmt.Sprintf("Player %d pays Player %d $%d", playerA.ID, playerB.ID, playerB.Money))
			playerB.Money = 0
			j--
		}
	}
	return payments
}

func displayPayments(payments []string) {
	if len(payments) > 0 {
		fmt.Println("\nPayments:")
		for _, payment := range payments {
			fmt.Println(payment)
		}
	} else {
		fmt.Println("\nNo payments required.")
	}
}
