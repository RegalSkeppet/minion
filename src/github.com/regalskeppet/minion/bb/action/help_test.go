package action

import "testing"

func TestHelpUnrecognized(t *testing.T) {
  a, err := GetAction("help")
  if err != nil {
    t.Fatal(err)
  }
  err = a([]string{"unrecognized"})
  if err == nil {
    t.Fatal("Should have gotten an error.")
  }
}

func ExampleHelp() {
  a, err := GetAction("help")
  if err != nil {
    return
  }
  err = a(nil)
  if err != nil {
    return
  }
  // Output:
  // Usage: bb <action> [options]
  //
  // Actions:
  //     help                  Show this help.
  //     apply [file]          Provision using the specified setup file.
}

func ExampleHelpApply() {
  a, err := GetAction("help")
  if err != nil {
    return
  }
  err = a([]string{"apply"})
  if err != nil {
    return
  }
  // Output:
  // Usage: bb apply [file]
  //
  // Provision using the specified setup file.
  //
  // File:
  //     If file is specified, use it as configuration.
  //     If file is missing, use stdin as configuration.
}

