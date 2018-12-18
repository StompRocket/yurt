package yurt

import (
    "strings"
    "regexp"
    "fmt"
)

type command struct {
    Expr regexp.Regexp
    PlainText string
    Desc string
    Handler func([]string)
}

type Parser struct {
    commands []command
}

func (p *Parser) Command(match string, desc string, handler func([]string)) {
    cmd := strings.Split(match, " ")
    expr := "^"
    // translate all this to regex
    for i, c := range cmd {
        if strings.HasPrefix(c, "$") {
            switch c {
            case "$s":
                expr += `(\w+)`
            case "$d":
                expr += `(\d+\.?\d*)`
            }
        } else {
            expr += c
        }
        if len(cmd)-1 != i {
            expr += " "
        }
    }
    expr += "$"

    p.commands = append(p.commands, command{
        Expr: *regexp.MustCompile(expr),
        PlainText: match,
        Desc: desc,
        Handler: handler,
    })
}

func (p *Parser) Run(args []string) {
    argstr := strings.Join(args[1:], " ")
    for _, cmd := range p.commands {
        matches := cmd.Expr.FindAllString(argstr, -1)
        if matches != nil {
            cmd.Handler(matches)
            break
        }
    }
}

func (p *Parser) Help() {
    for _, cmd := range p.commands {
        if strings.Replace(cmd.PlainText, " ", "", -1) != "" {
            fmt.Printf("%-15s %s\n", cmd.PlainText, cmd.Desc)
        }
    }
}