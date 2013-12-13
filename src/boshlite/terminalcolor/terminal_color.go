package terminalcolor

import "fmt"

type Color uint

const (
  Red     Color = 31
  Green         = 32
  Yellow        = 33
  Blue          = 34
  Magenta       = 35
  Cyan          = 36
)

func Colorize(message string, color Color, bold bool) string {
  attr := 0
  if bold {
    attr = 1
  }

  return fmt.Sprintf("\033[%d;%dm%s\033[0m", attr, color, message)
}

func FailureColor(message string) string {
  return Colorize(message, Red, false)
}

func SuccessColor(message string) string {
  return Colorize(message, Green, false)
}

func WarnColor(message string) string {
  return Colorize(message, Yellow, false)
}
