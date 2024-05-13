package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

const BOARD_SIZE = 3
const ZERO_VALUE = 2
const TIE_GAME = 0

var runningCode uint = ZERO_VALUE

type Player struct {
	label string
	value uint
}

var player Player = Player{"X", 1}
var enemy Player = Player{"O", 3}

func printBoard(board [BOARD_SIZE][BOARD_SIZE]uint) {

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			value := "-"
			if board[i][j] == player.value {
				value = "X"
			} else if board[i][j] == BOARD_SIZE {
				value = "O"
			}
			fmt.Printf("%s ", value)
		}
		fmt.Println()
	}
	fmt.Println()
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func boardIsFull(board [BOARD_SIZE][BOARD_SIZE]uint) bool {

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if board[i][j] == ZERO_VALUE {
				return false
			}
		}
	}
	return true

}

func MainGame(board *[BOARD_SIZE][BOARD_SIZE]uint) {

	printBoard(*board)

	fmt.Println("Choose a position (line, column):")
	var i, j uint
	nItens, err := fmt.Scanln(&i, &j)

	if err != nil {
		fmt.Println(err)
		return
	}

	if nItens != 2 {
		fmt.Println("Invalid number of arguments")
		return
	}

	if i >= BOARD_SIZE || j >= BOARD_SIZE {
		fmt.Println("Invalid position")
		return
	}

	if board[i][j] != ZERO_VALUE {
		fmt.Println("Position already taken")
		return
	}

	clearScreen()

	board[i][j] = player.value

	runningCode = verifyWinning(*board)

	if runningCode != ZERO_VALUE {
		return
	}
	enemyPlays(board)

	runningCode = verifyWinning(*board)
}

func enemyPlays(board *[BOARD_SIZE][BOARD_SIZE]uint) {
	playerAlmostWin := BOARD_SIZE*player.value + 1
	enemyAlmostWin := BOARD_SIZE*enemy.value - 2
	for i := 0; i < BOARD_SIZE; i++ {
		sumLine := reduce(board[i], func(acc, current uint) uint {
			return acc + current
		}, 0)

		arrayColumn := [BOARD_SIZE]uint{board[0][i], board[1][i], board[2][i]}
		sumColumn := reduce(arrayColumn, func(acc, current uint) uint {
			return acc + current
		}, 0)

		if sumLine == enemyAlmostWin || sumColumn == enemyAlmostWin || sumLine == playerAlmostWin || sumColumn == playerAlmostWin {
			index := findIndex(board[i], ZERO_VALUE)

			if index != -1 {
				board[i][index] = enemy.value
				return
			}

		}
	}

	diagonalLeft := [BOARD_SIZE]uint{board[0][0], board[1][1], board[2][2]}
	diagonalRight := [BOARD_SIZE]uint{board[0][2], board[1][1], board[2][0]}

	sumDiagonalLeft := reduce(diagonalLeft, func(acc, current uint) uint {
		return acc + current
	}, 0)

	sumDiagonalRight := reduce(diagonalRight, func(acc, current uint) uint {
		return acc + current
	}, 0)

	if sumDiagonalLeft == enemyAlmostWin || sumDiagonalRight == enemyAlmostWin || sumDiagonalLeft == playerAlmostWin || sumDiagonalRight == playerAlmostWin {

		index := findIndex(diagonalLeft, ZERO_VALUE)
		if index != -1 {
			board[index][index] = enemy.value
			return
		}

		index = findIndex(diagonalRight, ZERO_VALUE)
		if index != -1 {
			board[index][2-index] = enemy.value
			return
		}
	}

	for {
		i := rand.Intn(3)
		j := rand.Intn(3)
		if board[i][j] == ZERO_VALUE {
			board[i][j] = enemy.value
			break
		}
	}

}

func reduce(s [BOARD_SIZE]uint, f func(uint, uint) uint, initValue uint) uint {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func findIndex(arr [3]uint, target uint) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func verifyWinning(board [BOARD_SIZE][BOARD_SIZE]uint) uint {
	playerWin := BOARD_SIZE * player.value
	enemyWin := BOARD_SIZE * enemy.value

	for i := 0; i < BOARD_SIZE; i++ {
		sumLine := reduce(board[i], func(acc, current uint) uint {
			return acc + current
		}, 0)

		arrayColumn := [BOARD_SIZE]uint{board[0][i], board[1][i], board[2][i]}
		sumColumn := reduce(arrayColumn, func(acc, current uint) uint {
			return acc + current
		}, 0)

		if sumLine == playerWin || sumColumn == playerWin {
			return player.value
		} else if sumLine == enemyWin || sumColumn == enemyWin {
			return enemy.value
		}
	}

	sumDiagonalLeft := reduce([BOARD_SIZE]uint{board[0][0], board[1][1], board[2][2]}, func(acc, current uint) uint {
		return acc + current
	}, 0)

	sumDiagonalRight := reduce([BOARD_SIZE]uint{board[0][2], board[1][1], board[2][0]}, func(acc, current uint) uint {
		return acc + current
	}, 0)

	if sumDiagonalLeft == playerWin || sumDiagonalRight == playerWin {
		return player.value
	} else if sumDiagonalLeft == enemyWin || sumDiagonalRight == enemyWin {
		return enemy.value
	}

	boardIsFull := boardIsFull(board)
	if boardIsFull {
		return TIE_GAME
	}

	return ZERO_VALUE

}

func main() {
	board := [BOARD_SIZE][BOARD_SIZE]uint{
		{ZERO_VALUE, ZERO_VALUE, ZERO_VALUE},
		{ZERO_VALUE, ZERO_VALUE, ZERO_VALUE},
		{ZERO_VALUE, ZERO_VALUE, ZERO_VALUE},
	}
	for runningCode == ZERO_VALUE {
		MainGame(&board)
	}
	printBoard(board)

	if runningCode == TIE_GAME {
		fmt.Println("It's a tie!")
		return
	}

	var winner Player
	if runningCode == player.value {
		winner = player
	} else {
		winner = enemy
	}
	fmt.Println("The winner is ", winner.label)

}
