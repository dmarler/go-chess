package main

import "github.com/dmarler/go-chess/chess"

func main() {
  board := chess.ParseFENString("rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b KQkq c3 0 2")

  board.PrintBoard()
}
