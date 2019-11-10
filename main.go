// Two-player Go Command Line Tic Tac Toe
/*
  TODO:
  Computerized AI
*/

package main

import "fmt"
import "strconv"
import "regexp"

var validNumber = regexp.MustCompile("^[1-9]{1}$")
type Board [3][3]string

func createBoard() Board {
  x := 0
  var b Board
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      x++
      b[i][j] = strconv.Itoa(x)
    }
  }
  return b
}

func printBoard(b Board) {
  fmt.Printf("\t%s|%s|%s\n\t%s|%s|%s\n\t%s|%s|%s\n", b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2], b[2][0], b[2][1], b[2][2])
}

func markBoard(b Board, player int, input string) Board {
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if b[i][j] == input {
        switch player {
          case 1: b[i][j] = "X"
          case 2: b[i][j] = "O"
        }
      }
    }
  }
  return b
}

func checkWinner(b Board, player int) bool {
  var mark string
  switch player {
    case 1: mark = "X"
    case 2: mark = "O"
  }
  for i := 0; i < 3; i++ {
    horizontal := 0
    vertical := 0
    for j := 0; j < 3; j++ {
      if b[i][j] == mark {
        horizontal++
      }
      if b[j][i] == mark {
        vertical++
      }
      if horizontal == 3 || vertical == 3 {
        return true
      }
    }
  }
  // Check Diagonal
  if b[0][0] == mark && b[1][1] == mark && b[2][2] == mark {
    return true
  }
  if b[2][0] == mark && b[1][1] == mark && b[0][2] == mark {
    return true
  }
  return false
}

func getInput(b Board, player int) (string, bool) { 
  fmt.Printf("Player %d: Make your move, select a # 1-9 (or 'exit')\n", player)
  var input string
  fmt.Scanf("%s", &input)
  if !validNumber.MatchString(input) {
    fmt.Println("Invalid input.  Please try again.")
    return input, false
  }
  return input, true
}

func mainLoop() {
  gameOver := false 
  player := 1
  b := createBoard()
  moves := 0
  for gameOver == false {
    moves++
    validInput := false
    var input string
    for validInput == false {
      input, validInput = getInput(b, player)
      b = markBoard(b, player, input)
    }
    printBoard(b)
    if checkWinner(b, player) == true {
      fmt.Printf("Player %d wins!\n", player)
      gameOver = true
      break
    }
    switch player {
      case 1: player = 2
      case 2: player = 1
    }
    if input == "exit" {
      fmt.Println("Game over")
      gameOver = true
    }
    if moves == 9 {
      fmt.Println("It's a tie!")
      gameOver = true
    }
  }
}

func main() {
  fmt.Println("Two-player Go Command line Tic-tac-toe")
  b := createBoard()
  printBoard(b)
  mainLoop()
}
