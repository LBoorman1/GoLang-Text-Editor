package text

import (
  "fmt"
)

func MoveCursorRight(x int) {
  fmt.Printf("\x1b[%dC", x)
}

func MoveCursorLeft(x int) {
  fmt.Printf("\x1b[%dD", x)
}

func MoveCursorUp(x int) {
  fmt.Printf("\x1b[%dA", x)
}

func MoveCursorDown(x int) {
  fmt.Printf("\x1b[%dB", x)
}
