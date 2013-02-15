package action

import (
  "testing"
)

func TestApplyWithoutFileAndClosedStdin(t *testing.T) {
  a, err := GetAction("apply")
  if err != nil {
    t.Fatal(err)
  }
  err = a(nil)
  if err == nil {
    t.Fatal(err)
  }
}
