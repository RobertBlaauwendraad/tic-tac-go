package main

import "fmt"

func main() {
	// Create a 3x3 array that represents our tic tac toe board
	var board [3][3]string

	// Initialize our board with dashes (empty positions)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}

	// Create a Scanner and ask the players for their names
	fmt.Println("Let's play Tic Tac Go!")
	fmt.Print("Player 1, what is your name? ")
	var p1 string
	fmt.Scanln(&p1)
	fmt.Print("Player 2, what is your name? ")
	var p2 string
	fmt.Scanln(&p2)

	// Create a player1 boolean that is true if it is player 1's turn and false if it is player 2's turn
	var player1 bool = true

	// Create a gameEnded boolean and use it as the condition in the while loop
	var gameEnded bool = false

	for !gameEnded {
		// Draw the board
		drawBoard(board)

		// Print whose turn it is
		if player1 {
			fmt.Println(p1 + "'s Turn (x):")
		} else {
			fmt.Println(p2 + "'s Turn (o):")
		}

		// Create a char variable that stores either 'x' or 'o' based on what player's turn it is
		var c string = "-"
		if player1 {
			c = "x"
		} else {
			c = "0"
		}

		// Create row and col variables which represent indexes that correspond to a position on our board
		var row int
		var col int

		// Only break out of the while loop once the user enters a valid position
		for {
			fmt.Print("Enter a row number (0, 1, or 2): ")
			fmt.Scanln(&row)
			fmt.Print("Enter a column number (0, 1, or 2): ")
			fmt.Scanln(&col)

			// Check if the row and col are 0, 1, or 2
			if row < 0 || col < 0 || row > 2 || col > 2 {
				fmt.Println("This position is off the bounds of the board! Try again.")
				continue

				// Check if the position on the board the user entered is empty (has a -) or not
			} else if board[row][col] != "-" {
				fmt.Println("Someone has already made a move at this position! Try again.")
				continue

				// Otherwise, the position is valid so break out of the while loop
			} else {
				break
			}
		}

		// Set the postion on the board at row, col to c
		board[row][col] = c

		// Check to see if either play has won
		if playerHasWon(board) == "x" {
			fmt.Println(p1 + " has won!")
			gameEnded = true
		} else if playerHasWon(board) == "0" {
			fmt.Println(p2 + " has won!")
			gameEnded = true
		} else {
			// If neither player has won, check to see if there has been a tie (if the board is full)
			if boardIsFull(board) {
				fmt.Println("It's a tie!")
				gameEnded = true
			} else {
				// If player1 is true, make it false, and vice versa; this way, the players alternate each turn
				player1 = !player1
			}
		}
	}

	// Draw the board at the end of the game
	drawBoard(board)
}

// Make a function to draw the tic tac toe board
func drawBoard(board [3][3]string) {
	fmt.Println("Board:")
	for i := 0; i < 3; i++ {
		// The inner for loop prints out each row of the board
		for j := 0; j < 3; j++ {
			fmt.Print(board[i][j])
		}
		// This print statement makes a new line so that each row is on a separate line
		fmt.Println()
	}
}

// Make a function to see if someone has won and return the winning char
func playerHasWon(board [3][3]string) string {
	// Check each row
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][0] != "-" {
			return board[i][0]
		}
	}

	// Check each column
	for j := 0; j < 3; j++ {
		if board[0][j] == board[1][j] && board[1][j] == board[2][j] && board[0][j] != "-" {
			return board[0][j]
		}
	}

	// Check the diagnols
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[0][0] != "-" {
		return board[0][0]
	}
	if board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[2][0] != "-" {
		return board[2][0]
	}

	// Otherwise nobody has not won yet
	return " "
}

// Make a function to check if all of the positions on the board have been filled
func boardIsFull(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "-" {
				return false
			}
		}
	}
	return true
}
