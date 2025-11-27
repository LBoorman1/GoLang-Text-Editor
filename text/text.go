// Package text provides commonly used sequences for manipulating the terminal
// ANSII escape sequences are held here for text manipulation.
package text

import (
  "fmt"
)

var SpecialChars = map[int]func([]byte) []byte{
  8: DeleteLastChar,
  127: DeleteLastChar,
}

// ClearScreen clears all characters from the terminal and returns the cursor HOME
func ClearScreen() {
    fmt.Print("\x1b[2J\x1b[H")
}

func DeleteLastChar(b []byte) []byte {
  if len(b) > 0 {
    b = b[:len(b)-1]
    return b
  }
  return nil
}

// HandleChar accepts incoming characters and handles them accordingly
func HandleChar(buf []byte, b byte) []byte {
  if fn, ok := SpecialChars[int(b)]; ok {
    buf = fn(buf)
  } else {
    buf = append(buf, b)
  }
  return buf
}
