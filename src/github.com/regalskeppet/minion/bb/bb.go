package main

import (
  "os"
  "fmt"
  "github.com/regalskeppet/minion/bb/action"
)

func main() {
  arg0 := "help"
  args := os.Args[1:]

  if len(os.Args) >= 2 {
    arg0 = os.Args[1]
    args = os.Args[2:]
  }

  action, err := action.GetAction(arg0)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  err = action(args)

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
