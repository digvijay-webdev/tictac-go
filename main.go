package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CLEAR_SC() {
	if runtime.GOOS == "darwin" {
		sc := exec.Command("clear")
		sc.Stdout = os.Stdout
		sc.Run()
	} else {
		sc := exec.Command("cls")
		sc.Stdout = os.Stdout
		sc.Run()
	}
}

func SELECT_PLAYERS() (string, string) {
	var (
		player_one string
		player_two string
	)

	for {
		fmt.Println("\nPlayer One (Enter Name In One Word):")
		fmt.Scanln(&player_one)

		if player_one != "" {
			break
		} else {
			continue
		}
	}

	for {
		fmt.Println("\nPlayer Two (Enter Name In One Word):")
		fmt.Scanln(&player_two)

		if player_two != "" {
			break
		} else {
			continue
		}
	}

	CLEAR_SC()

	return player_one, player_two
}

func DRAW_BOARD(b [3][3]string) {
	for i1 := 0; i1 < len(b); i1++ {
		if i1 == 1 {
			fmt.Println("---------")
			fmt.Println(b[i1][0], "|", b[i1][1], "|", b[i1][2])
			fmt.Println("---------")
		} else {
			fmt.Println(b[i1][0], "|", b[i1][1], "|", b[i1][2])
		}
	}
}

func PROCESS_PLOTS(b [3][3]string, bi string, ap [9]string, symbol string) ([3][3]string, [9]string) {
	for i := 0; i < len(ap); i++ {
		if bi == ap[i] {
			switch bi {
			case "1":
				ap[0] = "taken"
				b[0][0] = symbol
			case "2":
				ap[1] = "taken"
				b[0][1] = symbol
			case "3":
				ap[2] = "taken"
				b[0][2] = symbol
			case "4":
				ap[3] = "taken"
				b[1][0] = symbol
			case "5":
				ap[4] = "taken"
				b[1][1] = symbol
			case "6":
				ap[5] = "taken"
				b[1][2] = symbol
			case "7":
				ap[6] = "taken"
				b[2][0] = symbol
			case "8":
				ap[7] = "taken"
				b[2][1] = symbol
			case "9":
				ap[8] = "taken"
				b[2][2] = symbol
			default:
				fmt.Println("Invalid input or plot taken..")
			}
		}
	}

	DRAW_BOARD(b)

	// board, available_plots
	return b, ap
}

/*
	win conditions:
		on x-axis:
			1,2,3 || 4,5,6 || 7,8,9
		on y-axis:
			1,4,7 || 2,5,8 || 3,6,9
		on diagonals:
			1,5,9 || 3,5,7
*/

func DETECT_WINNER(b [3][3]string, p1 string, p2 string) (r string) {
	var result string

	// check x-axis
	if (b[0][0] == "X" && b[0][1] == "X" && b[0][2] == "X") || (b[1][0] == "X" && b[1][1] == "X" && b[1][2] == "X") || (b[2][0] == "X" && b[2][1] == "X" && b[2][2] == "X") {
		result = p1 + " WON THE GAME!!"
	}
	if (b[0][0] == "O" && b[0][1] == "O" && b[0][2] == "O") || (b[1][0] == "O" && b[1][1] == "O" && b[1][2] == "O") || (b[2][0] == "O" && b[2][1] == "O" && b[2][2] == "O") {
		result = p2 + " WON THE GAME!!"
	}

	// check y-axis
	if (b[0][0] == "X" && b[1][0] == "X" && b[2][0] == "X") || (b[0][1] == "X" && b[1][1] == "X" && b[2][1] == "X") && (b[0][2] == "X" && b[1][2] == "X" && b[2][2] == "X") {
		result = p1 + " WON THE GAME!!"
	}
	if (b[0][0] == "O" && b[1][0] == "O" && b[2][0] == "O") || (b[0][1] == "O" && b[1][1] == "O" && b[2][1] == "O") && (b[0][2] == "O" && b[1][2] == "O" && b[2][2] == "O") {
		result = p2 + " WON THE GAME!!"
	}

	// check diagonals
	if (b[0][0] == "X" && b[1][1] == "X" && b[2][2] == "X") || (b[0][2] == "X" && b[1][1] == "X" && b[2][1] == "X") {
		result = p1 + " WON THE GAME!!"
	}
	if (b[0][0] == "O" && b[1][1] == "O" && b[2][2] == "O") || (b[0][2] == "O" && b[1][1] == "O" && b[2][1] == "O") {
		result = p2 + " WON THE GAME!!"
	}

	return result
}

func main() {
	fmt.Println("Welcome to Tic-Tac-Toe Game")

	// select initiate option
	fmt.Println("\nType 'exit' to quit the game or press 'enter' to begin:")
	var init_options string
	fmt.Scanln(&init_options)
	if init_options == "exit" || init_options == "Exit" || init_options == "EXIT" {
		// exit program
	} else {
		// process
		player_one, player_two := SELECT_PLAYERS()
		fmt.Println("Get Ready to play (", player_one, "and", player_two, ")")

		// draw board
		var board = [3][3]string{
			{"1", "2", "3"},
			{"4", "5", "6"},
			{"7", "8", "9"},
		}
		DRAW_BOARD(board)

		// take player inputs
		var board_input string
		var available_plots = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		var result string

		for {
			// taking input from player one
			for {
				fmt.Println(player_one + "'s turn:")
				fmt.Scanln(&board_input)

				if len(board_input) == 1 {
					board, available_plots = PROCESS_PLOTS(board, board_input, available_plots, "X")
					break
				} else {
					fmt.Println("Wrong Input, try again..")
				}
			}

			// print result
			result = DETECT_WINNER(board, player_one, player_two)
			if len(result) > 5 {
				fmt.Println(result)
				break
			}

			// taking input from player two
			for {
				fmt.Println(player_two + "'s turn:")
				fmt.Scanln(&board_input)

				if len(board_input) == 1 {
					board, available_plots = PROCESS_PLOTS(board, board_input, available_plots, "O")
					break
				} else {
					fmt.Println("Wrong Input, try again..")
				}
			}

			// print result
			result = DETECT_WINNER(board, player_one, player_two)
			if len(result) > 5 {
				fmt.Println(result)
				break
			}
		}
	}
}
