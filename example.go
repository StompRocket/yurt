package main

import (
    "github.com/StompRocket/colors"
    "github.com/StompRocket/yurt/box"
    "github.com/StompRocket/yurt"
    "fmt"
    "os"
)

func main() {
    var p yurt.Parser

    // empty strings == default
    p.Command("", "", func(_ []string) {
        p.Help()
    })

    p.Command("foo bar", "An example", func(matches []string) {
        fmt.Println(matches)
    })

    p.Command("baz quux", "Another example", func(matches []string) {
        fmt.Println(matches)
    })

    p.Command("-h", "Shows this message", func(_ []string) {
        box.Box("Help", colors.Cyan)
        fmt.Printf(`Example yurt program
%sgithub.com/StompRocket/yurt%s

`, colors.ColorString(colors.Cyan), colors.ColorString(colors.Reset))
        p.Help()
    })

    p.Run(os.Args)
}