package colors

import (
    "fmt"
)

type Color int

const (
    Reset Color = 0

    Black Color = iota + 29
    Red
    Green
    Yellow
    Blue
    Magenta
    Cyan
    White
    BBlack
    BRed
    BGreen
    BYellow
    BBlue
    BCyan
    BWhite
)

func ColorString(c Color) string {
    return fmt.Sprintf("\033[%dm", c)
}