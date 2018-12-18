package yurt

import (
    "strings"
    "regexp"
    "fmt"
)

// A private type used internally within the 
// Parser struct to store regex and plaintest data
type command struct {
    Expr regexp.Regexp
    PlainText string
    Desc string
    Handler func([]string)
}

type Parser struct {
    commands []command
}

// Adds a command that matches pseudo-regex `match` with description
// `desc`, and a lambda that takes `[]string` to handle if matched.
// Pseudo-regex allows use of sequences like `$s` and `$d` to more
// easily match things without having to use complex regex.
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

// Parses the `[]string` passed as a parameter and calls
// the appropriate command stored inside the Parser, passing
// the matches as a parameter to the lambda.
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

// Prints out a help message based on teh commands present in the parser.
func (p *Parser) Help() {
    for _, cmd := range p.commands {
        if strings.Replace(cmd.PlainText, " ", "", -1) != "" {
            fmt.Printf("%-15s %s\n", cmd.PlainText, cmd.Desc)
        }
    }
}