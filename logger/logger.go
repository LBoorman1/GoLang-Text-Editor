package logger

import (
  "log"
  "os"
)

var Log = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func Info(msg string) { Log.SetPrefix("INFO: "); Log.Output(2, msg) }
func Debug(msg string) { Log.SetPrefix("DEBUG: "); Log.Output(2, msg) }
func Critical(msg string) { Log.SetPrefix("CRITICAL: "); Log.Output(2, msg) }
