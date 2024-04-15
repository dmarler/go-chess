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
