package text

import (
	"syscall"
	"unsafe"
	"fyle.com/text/logger"
)

type termios syscall.Termios

// GetTerminalState gets the current terminal state from the syscall.IOCTL call
// using TCGETS "Gets" the current terminal state and sets it
func GetTerminalState(fd uintptr) (*termios, error) {
  var oldState termios
  
  _, _, err := syscall.Syscall(
    syscall.SYS_IOCTL, 
    fd, 
    uintptr(syscall.TCGETS), 
    uintptr(unsafe.Pointer(&oldState)),
  )
  if err != 0 {
    logger.Critical("Could not get current term state")
    return nil, err
  }
  return &oldState, nil
}

func SetTerminalState(fd uintptr, newState *termios) error {
  _, _, err := syscall.Syscall(
    syscall.SYS_IOCTL,
    fd,
    uintptr(syscall.TCSETS),
    uintptr(unsafe.Pointer(newState)),
  )
  if err != 0 {
    return err
  }
  return nil
}

func EnableRawMode(fd uintptr) (*termios, error) {
  oldState, err := GetTerminalState(fd)

  if err != nil {
    return nil, err
  }

  terminalRawState := *oldState
  terminalRawState.Lflag &^= syscall.ICANON | syscall.ECHO | syscall.IEXTEN
  terminalRawState.Iflag &^= syscall.IXON | syscall.ICRNL | syscall.BRKINT | syscall.INPCK | syscall.ISTRIP | syscall.PARMRK
  terminalRawState.Oflag &^= syscall.OPOST
  terminalRawState.Cflag &^= syscall.CS8

  if err := SetTerminalState(fd, &terminalRawState); err != nil {
    return nil, err
  }

  return oldState, nil
}


