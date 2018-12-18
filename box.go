package yurt

import (
    "strings"
    "fmt"
)

const (
    topL = "╭"
    topR = "╮"
    btmL = "╰"
    btmR = "╯"
    up   = "│"
    side = "─"
)

// Prints an ASCII box to the terminal with contents
// `content` and a color string which will be prepended
// and appended to the output. It is recommended you use
// one of the color constants provided although it will
// not complain if you pass it something else. 
func Box(content string, color string) {
    length := len(content)

    bar := strings.Repeat(side, length + 2)

    // this is gonna be fun to maintain
    fmt.Printf(`%s%s%s%s
%s %s %s
%s%s%s%s
`,  color,
    topL, bar, topR, 
    up, content, up, 
    btmL, bar, btmR,
    Reset)
}