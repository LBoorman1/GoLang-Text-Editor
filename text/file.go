package text

import (
	"fmt"
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
  var newFileName []byte
  var b = make([]byte, 1)
  if fileName == "" {
    ClearScreen()
    fmt.Print("Enter file name: ")
    for {
      os.Stdin.Read(b)
      if int(b[0]) == 13 {
        break
      } 
      newFileName = HandleChar(newFileName, b[0])
      ClearScreen()
      fmt.Printf("Enter file name: %s", newFileName)
    }
    fileName = string(newFileName)
  }
  renderBuf = append(renderBuf, '\r', '\n')
  err := os.WriteFile(fileName, renderBuf, 0700)
  if err != nil {
    panic(err)
  }
}
