package lib

import (
  "io"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "encoding/base64"
  "os"
)

const helpText = `A File order will make sure a file has a certain content. Reads config from stdin:
  {
    "path": "/absolute/path/to/file",
    "content": "base64 encoded content"
  }
`

type fileOrder struct {
  Path string
  Content string
  Mode string
}

func Run(args []string, reader io.Reader) error {
  if len(args) > 1 && args[1] == "help" {
    fmt.Println(helpText)
    return nil
  }
  var order fileOrder
  dec := json.NewDecoder(reader)
  err := dec.Decode(&order)
  if err != nil {
    return err
  }
  content, err := base64.StdEncoding.DecodeString(order.Content)
  if err != nil {
    return err
  }
  var mode os.FileMode
  if order.Mode == "" {
    mode = 0644
  } else {
    _, err = fmt.Sscanf(order.Mode, "%o", &mode)
    if err != nil {
      return err
    }
  }
  err = ioutil.WriteFile(order.Path, content, mode)
  return err
}
