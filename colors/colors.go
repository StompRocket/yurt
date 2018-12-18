package colors

const (
    Reset string = "\033[0m"

    Black string = "\033[" + string(iota + 30) + "m"
    Red
    Green
    Yellow
    Blue
    Magenta
    Cyan
    White
)

const (
    BBlack string = "\033[" + string(iota + 30) + ";1m"
    BRed
    BGreen
    BYellow
    BBlue
    BCyan
    BWhite
)