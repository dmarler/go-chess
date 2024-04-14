package chess

import (
	"fmt"
	"strings"
)

type Board struct {
  squares [8][8]int
  active bool // false = white, true = black
  turn int
  castling int8
}

const (
  castleWhiteKing = 0b0001
  castleWhiteQueen = 0b0010
  castleBlackKing = 0b0100
  castleBlackQueen = 0b1000
)

const (
  Blank = iota
  BlackPawn
  BlackRook
  BlackBishop
  BlackKnight
  BlackQueen
  BlackKing
  WhitePawn
  WhiteRook
  WhiteBishop
  WhiteKnight
  WhiteQueen
  WhiteKing
)

func (b *Board) PrintBoard() {
  for x := 0; x < len(b.squares); x++ {
    for y := 0; y < len(b.squares[x]); y++ {
      fmt.Printf("%v ", b.squares[x][y])
    }
    fmt.Println("")
  }

  if !b.active {
    fmt.Println("Active: White")
  } else {
    fmt.Println("Active: Black")
  }
}

func (b *Board) GetSquare(i int) *int {
  return &b.squares[i/8][i%8]
}

func ParseFENString(fen string) Board {
  var b Board

  data := strings.Split(fen, " ")
  
  parseFenBoard(&b, data[0])

  return b
}

func parseFenBoard(b *Board, data string) {
  i := 0
  for _, c := range data {
    switch c {
    case 'r':
      *b.GetSquare(i) = BlackRook
      i++
    case 'n':
      *b.GetSquare(i) = BlackKnight
      i++
    case 'b':
      *b.GetSquare(i) = BlackBishop
      i++
    case 'q':
      *b.GetSquare(i) = BlackQueen
      i++
    case 'k':
      *b.GetSquare(i) = BlackKing
      i++
    case 'R':
      *b.GetSquare(i) = WhiteRook
      i++
    case 'N':
      *b.GetSquare(i) = WhiteKnight
      i++
    case 'B':
      *b.GetSquare(i) = WhiteBishop
      i++
    case 'Q':
      *b.GetSquare(i) = WhiteQueen
      i++
    case 'K':
      *b.GetSquare(i) = WhiteKing
      i++
    case 'p':
      *b.GetSquare(i) = BlackPawn
      i++
    case 'P':
      *b.GetSquare(i) = WhitePawn
      i++
    case '/':
      continue
    default:
      i += int(c - '0')
    }
  }
}
