package orders

import (
  "encoding/json"
  "regexp"
  "os"
  "io/ioutil"
  "encoding/base64"
  "strings"
)

type Order struct {
  Order string
  Details json.RawMessage
}

func ParseOrders(raw string) ([]*Order, error) {
  var orders []*Order
  err := json.Unmarshal([]byte(raw), &orders)
  if err != nil {
    return nil, err
  }
  reg, err := regexp.Compile("{\\s*\"encode_content\"\\s*:\\s*\"([^\"]+)\"\\s*}")
  if err != nil {
    return nil, err
  }
  for _, o := range orders {
    details := string(o.Details)
    for match := reg.FindStringSubmatch(details); match != nil; match = reg.FindStringSubmatch(details) {
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
      details = strings.Replace(details, match[0], encoded, -1)
    }
    o.Details = json.RawMessage(details)
  }
  return orders, err
}
