package chess

import (
	"testing"
)

func TestAlgToCord(t *testing.T) {
  col, row, err := AlgToCord("a7")
  if err != nil {
    t.Fatal(err)
  }
  if (col != 0 || row != 1) {
    t.Fatalf("Expected 'a7' to map to 0, 1 - Got %d, %d", col, row)
  }

  col, row, err = AlgToCord("h6")
  if err != nil {
    t.Fatal(err)
  }
  if (col != 7 || row != 2) {
    t.Fatalf("Expected 'a7' to map to 0, 1 - Got %d, %d", col, row)
  }

  _, _, err = AlgToCord("a66")
  if err == nil {
    t.Fatalf("Input 'a66' expected an error.")
  }
}

func TestMartialFEN(t *testing.T) {
  fens := []string{
    "8/5k2/3p4/1p1Pp2p/pP2Pp1P/P4P1K/8/8 b - - 99 50", 
    "rnbqkbnr/ppp1pppp/8/3p4/2PP4/8/PP2PPPP/RNBQKBNR b KQkq c3 0 2",
  }
  
  for _, fen := range fens {
    board := ParseFENString(fen)
    mfen := MarshallFENString(board)
    if mfen != fen {
      t.Fatalf("Incorrect FEN string marshalled:\nInput: %s\nOutput: %s", fen, mfen)
    }
  }
}
