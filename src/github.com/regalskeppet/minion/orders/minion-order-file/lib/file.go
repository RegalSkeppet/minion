package lib

import (
  "encoding/json"
  "os"
  "encoding/base64"
  "io/ioutil"
  "io"
  "fmt"
)

type fileOrderDescription struct {
  Path string
  Ensure string
  Content string
  ContentFile string
}

type fileOrder struct {
  Path string
  Content string
}

func decodeFileOrderDescription(reader io.Reader) (fileOrderDescription, error) {
  desc := fileOrderDescription{}
  dec := json.NewDecoder(reader)
  err := dec.Decode(&desc)
  return desc, err
}

func createFileOrder(desc fileOrderDescription) (*fileOrder, error) {
  file, err := os.Open(desc.ContentFile)
  if err != nil {
    return nil, err
  }
  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    return nil, err
  }
  order := new(fileOrder)
  order.Path = desc.Path
  order.Content = base64.StdEncoding.EncodeToString(bytes)
  return order, nil
}

func Run(args []string, reader io.Reader) {
  if len(args) > 2 && args[2] == "help" {
    fmt.Println(`Order: File
      {
        "path": "/absolute/path/to/file",
        "content": "simple content"
      }

      "path": Is the path of the file on the target system
      "ensure": Can be "present" (to make sure the file exists)
                      or
                      "absent" (to make sure the file doesn't exist)
      "content": Should be use if the content of the file is simple (a single row with allowed characters)
    `)
  }
}

/*func ApplyFile(conf *FileConfig) int64 {
  file, err := os.Create(conf.Path);
  if err != nil {
    log.Fatal(err)
  }
  content, err := base64.StdEncoding.DecodeString(conf.ContentFile)
  if err != nil {
    log.Fatal(err)
  }
  b, err := file.Write(content)
  if err != nil {
    log.Fatal(err)
  }
  return int64(b)
}*/
