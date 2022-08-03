package main

import (
	"fmt"
)

func main() {
	board := createBoard()
	winner, draw := play(board)
	if draw {
		fmt.Println("The game has ended in a draw")
	} else {
		fmt.Printf("\nPlayer %v has won the game!\n", winner)
	}
}

func play(board [3][3]string) (int, bool) {
	playerPieces := [2]string{"X", "O"}
	player := 0
	for turn := 0; turn < 9; turn++ {
		printBoard(board)
		x, y := getPlayerInput(player, board)
		board[y][x] = playerPieces[player]
		if isWinner(board) {
			return player + 1, false
		}
		player = (player + 1) % 2
	}
	return 0, true
}

func isWinner(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		if isThreeInARow(board[i][0], board[i][1], board[i][2]) || isThreeInARow(board[0][i], board[1][i], board[2][i]) {
			return true
		}
	}
	return isThreeInARow(board[0][0], board[1][1], board[2][2]) || isThreeInARow(board[2][0], board[1][1], board[0][2])
}

func isThreeInARow(space1 string, space2 string, space3 string) bool {
	return space1 == space2 && space2 == space3 && space3 != " "
}

func getPlayerInput(player int, board [3][3]string) (int, int) {
	var x, y int
	for {
		fmt.Printf("Player %d enter the coords of your next play ('x y' each should be 1, 2, or 3): ", player+1)
		fmt.Scan(&x, &y)
		if board[y][x] == " " {
			return x, y
		}

		fmt.Println("That space is already taken, pick another")
	}
}

func createBoard() [3][3]string {
	var board [3][3]string

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			board[y][x] = " "
		}
	}

	return board
}

func printBoard(board [3][3]string) {
	for y := 0; y < 3; y++ {
		fmt.Printf(" %v | %v | %v \n", board[y][0], board[y][1], board[y][2])
		if y < 2 {
			fmt.Println("---+---+---")
		}
	}
}
