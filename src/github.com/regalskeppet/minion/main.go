package main

import (
  "fmt"
  "github.com/regalskeppet/minion/action/apply"
)

func main() {
  conf := apply.Decode(`{
    "path": "/tmp/testfile2",
    "contentfile": "/tmp/testfile"
  }`)

  fmt.Println(conf)

  apply.ApplyFile(conf)
}
