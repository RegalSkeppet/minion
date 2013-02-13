package apply

import (
  "encoding/json"
  "strings"
  "os"
  "log"
  "encoding/base64"
  "io/ioutil"
)

type FileConfig struct {
  Path string
  ContentFile string
}

func Decode(raw string) *FileConfig {
  conf := new(FileConfig)
  dec := json.NewDecoder(strings.NewReader(raw))
  err := dec.Decode(conf)
  if err != nil {
    log.Fatal(err)
  }
  file, err := os.Open(conf.ContentFile)
  if err != nil {
    log.Fatal(err)
  }
  bytes, err := ioutil.ReadAll(file)
  if err != nil {
    log.Fatal(err)
  }
  conf.ContentFile = base64.StdEncoding.EncodeToString(bytes)
  return conf
}

func ApplyFile(conf *FileConfig) int64 {
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
}
