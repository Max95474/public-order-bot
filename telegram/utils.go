package telegram

import (
  "strings"
)

func ejectCommand(text, botName string) string {
  var command string
  
  command = strings.Replace(text, "@" + botName, "", 1)
  command = strings.Replace(command, "/", "", 1)

  return command
}