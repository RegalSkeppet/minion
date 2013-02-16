package orders

import (
  "encoding/json"
  "regexp"
  "os"
  "io/ioutil"
  "encoding/base64"
  "strings"
  "fmt"
)

type Order struct {
  Order string
  Details json.RawMessage
}

func ParseOrders(raw string) ([]*Order, error) {
  reg, err := regexp.Compile("{\\s*\"encode_content\"\\s*:\\s*\"([^\"]+)\"\\s*}")
  if err != nil {
    return nil, err
  }
  for match := reg.FindStringSubmatch(raw); match != nil; match = reg.FindStringSubmatch(raw) {
    file, err := os.Open(match[1][7:])
    if err != nil {
      return nil, err
    }
    content, err := ioutil.ReadAll(file)
    if err != nil {
      file.Close()
      return nil, err
    }
    err = file.Close()
    if err != nil {
      return nil, err
    }
    encoded := base64.StdEncoding.EncodeToString(content)
    raw = strings.Replace(raw, match[0], fmt.Sprintf("\"%s\"",  encoded), -1)
  }
  var orders []*Order
  err = json.Unmarshal([]byte(raw), &orders)
  return orders, err
}
