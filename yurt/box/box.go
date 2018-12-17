package box

import (
    "strings"
    "fmt"
    "../colors"
)

const (
    topL = "╭"
    topR = "╮"
    btmL = "╰"
    btmR = "╯"
    up   = "│"
    side = "─"
)

func Box(content string, color colors.Color) {
    length := len(content)

    bar := strings.Repeat(side, length + 2)

    // this is gonna be fun to maintain
    fmt.Printf(`%s%s%s%s
%s %s %s
%s%s%s%s
`,  colors.ColorString(color),
    topL, bar, topR, 
    up, content, up, 
    btmL, bar, btmR,
    colors.ColorString(colors.Reset))
}