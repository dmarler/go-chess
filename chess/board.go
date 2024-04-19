package chess

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
  squares [8][8]int
  active bool // true = white, false = black
  turn int
  castling int8
  ep string
  halfMoves int8
  fullMoves int16
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
  fmt.Println("   A B C D E F G H")
  fmt.Println("------------------")
  for x := 0; x < len(b.squares); x++ {
    fmt.Printf("%d| ", 8-x)
    for y := 0; y < len(b.squares[x]); y++ {
      fmt.Printf("%s ", printPiece(b.squares[x][y]))
    }
    fmt.Println("")
  }

  if b.active {
    fmt.Println("Active: White")
  } else {
    fmt.Println("Active: Black")
  }

  fmt.Printf("Castling Ability: %b\n", b.castling)
  fmt.Printf("En Passant: %s\n", b.ep)
  fmt.Printf("Halfmove Clock: %d\n", b.halfMoves/2)
  fmt.Printf("Fullmove Number: %d\n", b.fullMoves)
}

func (b *Board) GetSquare(i int) *int {
  return &b.squares[i/8][i%8]
}

func (b *Board) MovePiece(input string) {
  if len(input) == 4 {
    oldx, oldy, err := AlgToCord(input[0:2])
    if err != nil {
      fmt.Errorf("Error: %x", err)
    }
    newx, newy, err := AlgToCord(input[2:4])
    if err != nil {
      fmt.Errorf("Error: %x", err)
    }

    if b.squares[newy][newx] != Blank {
      b.halfMoves = 
    }

    b.squares[newy][newx] = b.squares[oldy][oldx]
    b.squares[oldy][oldx] = Blank

    fmt.Printf("Previous Move: %s\n", input)

    if !b.active {
      b.fullMoves++
    }
    b.active = !b.active
  } 
}

func printPiece(piece int) string {
  switch piece {
  case BlackPawn: return "p"
  case BlackRook: return "r"
  case BlackBishop: return "b"
  case BlackKnight: return "n"
  case BlackQueen: return "q"
  case BlackKing: return "k"
  case WhitePawn: return "P"
  case WhiteRook: return "R"
  case WhiteBishop: return "B"
  case WhiteKnight: return "N"
  case WhiteQueen: return "Q"
  case WhiteKing: return "K"
  case Blank: return " "
  default: return "#"
  }
}

func ParseFENString(fen string) Board {
  var b Board

  data := strings.Split(fen, " ")
  
  parseFenBoard(&b, data[0])
  parseActiveColor(&b, data[1])
  parseCastlingAbility(&b, data[2])
  parseEnPassant(&b, data[3])
  parseHalfMoves(&b, data[4])
  parseFullMoves(&b, data[5])

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

func parseActiveColor(b *Board, data string) {
  b.active = data == "w"
}

func parseCastlingAbility(b *Board, data string) {
  for _, c := range data {
    switch c {
    case '-':
      return
    case 'K':
      b.castling |= castleWhiteKing
    case 'Q':
      b.castling |= castleWhiteQueen
    case 'k':
      b.castling |= castleBlackKing
    case 'q':
      b.castling |= castleBlackQueen
    }
  }
}

func parseEnPassant(b *Board, ep string) error {
  if len(ep) != 2 && len(ep) != 1 {
    return fmt.Errorf("Expected En Passant to be a length of 2 or 1, got %d", len(ep))
  }
  b.ep = ep
  return nil
}

func parseHalfMoves(b *Board, hm string) error {
  halfMoves, err := strconv.ParseInt(hm, 10, 8)
  if err != nil {
    return err
  }
  b.halfMoves = int8(halfMoves)
  return nil
}

func parseFullMoves(b *Board, fm string) error {
  fullMoves, err := strconv.ParseInt(fm, 10, 16)
  if err != nil {
    return err
  }
  b.fullMoves = int16(fullMoves)
  return nil
}

func AlgToCord(pos string) (col int8, row int8, err error) {
  if (len(pos) != 2) {
    return 0, 0, fmt.Errorf("expected a string of length 2")
  }
  alg := []rune(pos)
  if (alg[0] > 'h' || alg[0] < 'a') {
    return 0, 0, fmt.Errorf("expected col between 'a' and 'h', got %v", alg[0])
  }
  if (alg[1] > '8' || alg[1] < '1') {
    return 0, 0, fmt.Errorf("expected row between 1 and 8m got %v", alg[1])
  }


  return int8(alg[0]-'a'), int8('8'-alg[1]), nil
}

func CordToAlg(col int8, row int8) (alg string, err error) {
  return "", nil
}
