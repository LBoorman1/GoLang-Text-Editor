package text

import (
	"log"
	"os"
)

func LoadFile(fileName string) []byte {
  data, err := os.ReadFile(fileName)
  if err != nil {
    log.Fatal(err)
  }
  return data
}
