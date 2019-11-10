// Two-player Go Command Line Tic Tac Toe
/*
  TODO:
  Computerized AI
*/

package main

import (
  "fmt"
  . "github.com/logrusorgru/aurora"
  "strconv"
  "regexp"
  "os"
)

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
  fmt.Printf("\t%s|%s|%s\n\t%s|%s|%s\n\t%s|%s|%s\n", pickColor(b[0][0]), pickColor(b[0][1]), pickColor(b[0][2]), pickColor(b[1][0]), pickColor(b[1][1]), pickColor(b[1][2]), pickColor(b[2][0]), pickColor(b[2][1]), pickColor(b[2][2]))
}

func pickColor(s string) Value {
  switch s {
    case "X": return Red(s).Bold()
    case "O": return Magenta(s).Bold()
    default:  return Cyan(s).Bold()
  }
}

func markBoard(b Board, player int, input string) (Board, bool) {
  validInput := false
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if b[i][j] == input {
        validInput = true
        switch player {
          case 1: b[i][j] = "X"
          case 2: b[i][j] = "O"
        }
      }
    }
  }
  if validInput == false {
    fmt.Println("That space is already taken, try again.")
  }
  return b, validInput
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

func getInput(player int) string {
  input := ""

  for !validNumber.MatchString(input) {
    fmt.Printf("Player %d: Make your move, select a # 1-9 (or 'exit')\n", player)
    fmt.Scanf("%s", &input)
    if input == "exit" {
      fmt.Println("Game over")
      os.Exit(0)
    }
    if !validNumber.MatchString(input) {
      fmt.Println("Invalid input")
    }
  }
  return input
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
      input = getInput(player)
      b, validInput = markBoard(b, player, input)
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
