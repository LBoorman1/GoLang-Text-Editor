// Package text provides commonly used sequences for manipulating the terminal
// ANSII escape sequences are held here for text manipulation.
package text

import (
  "fmt"
)

var SpecialChars = map[int]func([]byte) []byte{
  8: DeleteLastChar,
  127: DeleteLastChar,
  13: HandleCarriageReturn,
}

// ClearScreen clears all characters from the terminal and returns the cursor HOME
func ClearScreen() {
    fmt.Print("\x1b[2J\x1b[H")
}

func DeleteLastChar(b []byte) []byte {
  if len(b) <= 0 { 
    return nil
  }
  if len(b) >= 2 && b[len(b)-2] == '\r' && b[len(b)-1] == '\n' {
    b = b[:len(b)-2]
  } else {
    b = b[:len(b)-1]
  }
  return b
}

func HandleCarriageReturn(b []byte) []byte {
  b = append(b, '\r', '\n')
  return b
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
