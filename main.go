// Two-player Go Command Line Tic Tac Toe
// Copyright 2019 Drew Miller
// millerdrew@gmail.com

package main

import "fmt"
import "strconv"

func createBoard() [3][3]string {
  x := 0
  var b [3][3]string
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      x++
      b[i][j] = strconv.Itoa(x)
    }
  }
  return b
}

func printBoard(b[3][3]string) {
  fmt.Printf("\t%s|%s|%s\n\t%s|%s|%s\n\t%s|%s|%s\n", b[0][0], b[0][1], b[0][2], b[1][0], b[1][1], b[1][2], b[2][0], b[2][1], b[2][2])
}

func markBoard(b[3][3]string, player int, input string) [3][3]string {
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

func mainLoop() {
  gameOver := 0
  player := 1
  b := createBoard()
  for gameOver == 0 {
    fmt.Printf("Player %d: Make your move, select a # 1-9 (or 'exit')\n", player)
    var input string
    fmt.Scanf("%s", &input)
    b = markBoard(b, player, input)
    printBoard(b)
    switch player {
      case 1: player = 2
      case 2: player = 1
    }
    if input == "exit" {
      fmt.Println("Game over")
      gameOver = 1
    }
  }
}

func main() {
  fmt.Println("Two-player Go Command line Tic-tac-toe")
  b := createBoard()
  printBoard(b)
  mainLoop()
}