package orders

import (
  "testing"
  "os"
)

func TestParse(t *testing.T) {
  return
  raw := `[{
    "order": "testorder",
    "details": {}
  }]`

  orders, err := ParseOrders(raw)
  if err != nil {
    t.Fatal(err)
  }
  if len(orders) != 1 {
    t.Fatal("Number of orders != 1.")
  }
  if orders[0].Order != "testorder" {
    t.Fatal("First order wasn't \"testorder\".")
  }
  if string(orders[0].Details) != "{}" {
    t.Fatal("First order did not have empty details")
  }
}

func TestParseEncodeContent(t *testing.T) {
  err := os.MkdirAll("/tmp/minion_tests", 0755)
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
  err = file.Sync()
  if err != nil {
    t.Fatal(err)
  }
  err = file.Close()
  if err != nil {
    t.Fatal(err)
  }
  raw := `[{
    "order": "testorder",
    "details": {
      "encode_content": "file:///tmp/minion_tests/src"
    }
  }]`
  orders, err := ParseOrders(raw)
  if err != nil {
    t.Fatal(err)
  }
  if string(orders[0].Details) != "eyJ0ZXN0aW5n" {
    t.Fatalf("Content should be \"eyJ0ZXN0aW5n\", was \"%s\".", orders[0].Details)
  }
  os.RemoveAll("/tmp/minion_tests")
}
