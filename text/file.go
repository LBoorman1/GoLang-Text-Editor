package text

import (
	"log"
	"os"
)
var FileOperations = map[int]func(string, []byte) {
  19: SaveFile,
}

func LoadFile(fileName string) []byte {
  data, err := os.ReadFile(fileName)
  if err != nil {
    log.Fatal(err)
  }
  return data
}

func SaveFile(fileName string, renderBuf []byte) {
  if fileName == "" {
    return
  }
  renderBuf = append(renderBuf, '\r', '\n')
  err := os.WriteFile(fileName, renderBuf, 0700)
  if err != nil {
    panic(err)
  }
}
