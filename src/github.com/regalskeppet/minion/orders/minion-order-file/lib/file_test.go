package lib

import (
  "testing"
  "io/ioutil"
  "encoding/base64"
  "fmt"
  "strings"
  "os"
  "path/filepath"
)

func ExampleHelp() {
  Run([]string{"minion-order-file", "help"}, nil)
  // Output:
  // A File order will make sure a file has a certain content. Reads config from stdin:
  //   {
  //     "path": "/absolute/path/to/file",
  //     "content": "base64 encoded content"
  //   }
}

func TestRun(t *testing.T) {
  dir, err := ioutil.TempDir("", "minion-order-file-test-")
  defer os.RemoveAll(dir)
  if err != nil {
    t.Fatal(err)
  }
  path := filepath.Join(dir, "file")
  content := []byte("testcontent")
  enc := base64.StdEncoding.EncodeToString(content)
  filemode := os.FileMode(0600)
  raw := fmt.Sprintf(`{
    "path": "%s",
    "content": "%s",
    "mode": "%o"
  }`, path, enc, filemode)
  err = Run([]string{"minion-order-file"}, strings.NewReader(raw))
  if err != nil {
    t.Fatal(err)
  }
  file, err := ioutil.ReadFile(path)
  if err != nil {
    t.Fatal(err)
  }
  if string(content) != string(file) {
    t.Fatal(fmt.Sprintf("Expected content %s, got %s.", content, file))
  }
  fi, err := os.Stat(path)
  if err != nil {
    t.Fatal(err)
  }
  mode := fi.Mode()
  if mode != filemode {
    t.Fatal(fmt.Sprintf("Expected filemode %o, got %o.", filemode, mode))
  }
}

func TestRunDefaultFileMode(t *testing.T) {
  dir, err := ioutil.TempDir("", "minion-order-file-test-")
  defer os.RemoveAll(dir)
  if err != nil {
    t.Fatal(err)
  }
  path := filepath.Join(dir, "file")
  content := []byte("testcontent")
  enc := base64.StdEncoding.EncodeToString(content)
  raw := fmt.Sprintf(`{
    "path": "%s",
    "content": "%s"
  }`, path, enc)
  err = Run([]string{"minion-order-file"}, strings.NewReader(raw))
  if err != nil {
    t.Fatal(err)
  }
  fi, err := os.Stat(path)
  if err != nil {
    t.Fatal(err)
  }
  mode := fi.Mode()
  if mode != 0644 {
    t.Fatal(fmt.Sprintf("Expected filemode 644, got %o.", mode))
  }
}
