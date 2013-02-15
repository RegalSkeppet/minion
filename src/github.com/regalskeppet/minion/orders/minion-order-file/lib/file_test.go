package lib

import (
  "testing"
  "os"
  "strings"
)

func TestDecodeFileOrderDescription(t *testing.T) {
  t.Parallel()
  raw := `{
    "Path": "/tmp/minion_tests/dst",
    "ContentFile": "/tmp/minion_tests/src"
  }`
  desc, err := decodeFileOrderDescription(strings.NewReader(raw))
  if err != nil {
    t.Fatal(err)
  }
  if desc.Path != "/tmp/minion_tests/dst" {
    t.Errorf("Expected desc.Path to be \"/tmp/minion_tests/dest\", was \"%s\".", desc.Path)
  }
  if desc.ContentFile != "/tmp/minion_tests/src" {
    t.Errorf("Expected desc.Path to be \"/tmp/minion_tests/src\", was \"%s\".", desc.Path)
  }
}

func TestCreateFileOrder(t *testing.T) {
  err := os.RemoveAll("/tmp/minion_tests")
  if err != nil {
    t.Fatal(err)
  }
  err = os.Mkdir("/tmp/minion_tests", 0755)
  if err != nil {
    t.Fatal(err)
  }
  file, err := os.Create("/tmp/minion_tests/src")
  if err != nil {
    t.Fatal(err)
  }
  _, err = file.WriteString("{\"testing")
  if err != nil {
    t.Fatal(err)
  }
  err = file.Close()
  if err != nil {
    t.Fatal(err)
  }
  raw := `{
    "Path": "/tmp/minion_tests/dst",
    "ContentFile": "/tmp/minion_tests/src"
  }`
  desc, err := decodeFileOrderDescription(strings.NewReader(raw))
  if err != nil {
    t.Fatal(err)
  }
  order, err := createFileOrder(desc)
  if err != nil {
    t.Fatal(err)
  }
  if order.Path != "/tmp/minion_tests/dst" {
    t.Errorf("Expected order.Path to be \"/tmp/minion_tests/dest\", was \"%s\".", order.Path)
  }
  if order.Content != "eyJ0ZXN0aW5n" {
    t.Errorf("Expected order.Path to be \"eyJ0ZXN0aW5n\", was \"%s\".", order.Content)
  }
  err = os.RemoveAll("/tmp/minion_tests")
  if err != nil {
    t.Fatal(err)
  }
}
