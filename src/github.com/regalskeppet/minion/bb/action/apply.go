package action

import (
  "errors"
  "os"
  "encoding/json"
  "fmt"
)

type Config struct {
  Test string
}

func Apply(args []string) error {
  var dec *json.Decoder
  if len(args) == 0 {
    dec = json.NewDecoder(os.Stdin)
    if dec == nil {
      return errors.New("Got no dec.")
    }
  }
  var config Config
  err := dec.Decode(&config)
  if err != nil {
    return err
  }
  fmt.Println(config.Test)
  return nil
}
