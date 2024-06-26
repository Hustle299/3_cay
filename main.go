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
	var numPlayers, betMoney, totalRounds int

	// Input: Nhap so nguoi choi va tien cuoc moi van
	fmt.Print("Enter the number of players: ")
	fmt.Scanln(&numPlayers)
	fmt.Print("Enter the betting money for each round: ")
	fmt.Scanln(&betMoney)

	// Array de chua so van thang moi nguoi
	wins := make([]int, numPlayers)

	// Input: tong so van thang moi nguoi
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter the total winning rounds for player %d: ", i+1)
		fmt.Scanln(&wins[i])
		totalRounds += wins[i]
	}

	// Tinh toan so tien thang cua moi nguoi
	players := make([]Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = Player{ID: i + 1, Money: (wins[i] * (numPlayers - 1) * betMoney) - ((totalRounds - wins[i]) * betMoney)}
	}

	// Display so tien cua moi nguoi truoc khi tien hanh trả nợ
	fmt.Println("Final money for each player:")
	for _, player := range players {
		fmt.Printf("Player %d: %d\n", player.ID, player.Money)
	}

	// Tinh toan tien no va tra
	// Em se sort slice sao cho nguoi nao no nhieu tien nhat dung dau
	sort.Slice(players, func(i, j int) bool {
		return players[i].Money < players[j].Money
	})

	//slice de chua nhung string ghi lai giao dich
	var payments []string

	//khi bat dau A se la nguoi choi dau slice, B la nguoi choi cuoi slice
	for i, j := 0, numPlayers-1; i < j; {
		playerA, playerB := &players[i], &players[j]
		if playerA.Money >= 0 {
			break
		}

		transfer := playerA.Money
		//Nếu người chơi A dùng hết tiền đang âm trả cho người chơi B mà B vẫn thiếu tiền được trả
		// thì i++ vì người A không còn nợ phải trả nữa
		if playerB.Money+transfer >= 0 {
			playerA.Money += transfer
			//số tiền người chơi B còn thiếu để được trả
			playerB.Money += transfer
			payments = append(payments, fmt.Sprintf("Player %d pays Player %d $%d", playerA.ID, playerB.ID, -transfer))
			i++
		} else {
			//Nếu không thì	người B sẽ được nhận đủ tiền thì j--
			//Người A sẽ còn nợ số tiền đang có - số tiền người B
			playerA.Money += playerB.Money
			payments = append(payments, fmt.Sprintf("Player %d pays Player %d $%d", playerA.ID, playerB.ID, playerB.Money))
			playerB.Money = 0
			j--
		}
	}

	// In ra giao dich
	if len(payments) > 0 {
		fmt.Println("\nPayments:")
		for _, payment := range payments {
			fmt.Println(payment)
		}
	} else {
		fmt.Println("\nNo payments required.")
	}
}
