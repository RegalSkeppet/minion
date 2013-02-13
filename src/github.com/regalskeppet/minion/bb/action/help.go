package action

import (
  "fmt"
  "errors"
)

func Help(args []string) error {
  if len(args) == 0 {
    fmt.Println("Usage: bb <action> [options]\n")
    fmt.Println("Actions:")
    fmt.Println("    help                  Show this help.")
    fmt.Println("    apply [file]          Provision using the specified setup file.")
    return nil
  }
  switch (args[0]) {
  case "apply":
    fmt.Println("Usage: bb apply [file]\n")
    fmt.Println("Provision using the specified setup file.\n")
    fmt.Println("File:")
    fmt.Println("    If file is specified, use it as configuration.")
    fmt.Println("    If file is missing, use stdin as configuration.")
    break
  default:
    return errors.New("No help for unrecognized action.")
  }
  return nil
}
