package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
  "fyle.com/text/text"
)


func main() {
  fd := os.Stdin.Fd()
  
  // Gets the stable terminal state to be reset to later and enables
  // raw mode so that key presses are streamed directly to the application
  oldState, err := text.EnableRawMode(fd)
  if err != nil {
    panic(err)
  }
  defer text.SetTerminalState(fd, oldState)
  
  // create signal channel, buffer size of one to prevent missing signal
  // if signal sent before goroutine below starts.
  // Notify puts the signals into the sigs channel when they occur
  sigs := make(chan os.Signal, 1)
  signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
 
  // goroutine, enables main function to keep running
  // whilst still checking for signals such as SIGINT and SIGTERM
  // <- blocks until sigs receives something.
  go func() {
    <-sigs
    text.SetTerminalState(fd, oldState)
    os.Exit(0)
  }() 

  var b = make([]byte, 1)
  var fileBuf []byte
  var renderBuf []byte

  if len(os.Args) == 2 {
    fileName := os.Args[1]
    fileBuf = text.LoadFile(fileName)
    for _, c := range fileBuf {
      if c == '\n' {
        renderBuf = append(renderBuf, '\r', '\n')
      } else {
        renderBuf = append(renderBuf, c)
      }
    }
  }

  text.ClearScreen()  
  // Main event loop, handles processing of key presses.
  if len(renderBuf) > 0 {
    fmt.Print(string(renderBuf))
  }
  for {
    os.Stdin.Read(b)
    renderBuf = text.HandleChar(renderBuf, b[0])
    text.ClearScreen()
    fmt.Print(string(renderBuf))
  }
}
