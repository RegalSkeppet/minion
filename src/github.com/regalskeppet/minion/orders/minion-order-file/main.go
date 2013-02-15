package main

import (
  "os"
  "github.com/regalskeppet/minion/orders/minion-order-file/lib"
)

func main() {
  lib.Run(os.Args, os.Stdin)
}

