====
Yurt
====

A WIP Go CLI library

Installation
------------

::

    go get github.com/StompRocket/yurt

Usage
-----

This is pretty much the simplest program you can write with yurt. It parses the args
and does stuff with them. It's pretty self explanatory but there's more detail in the
comments

.. code-block :: go

    package main

    import (
        "github.com/StompRocket/yurt"
        "fmt"
        "os"
    )

    func main() {
        // Commands are stored in this
        var p yurt.Parser

        // Args: command string, description string, handler func([]string)
        // empty command is used as default
        p.Command("", "", func(_ []string) {
            // displays a help message showing all available commands
            p.Help()
        })

        // An actual command that actually works and is an actual command
        p.Command("foo", "An example", func(matches []string) {
            fmt.Println("Matches", matches)
        })

        // actually parse the args, and call the appropriate lambda
        p.Run(os.Args)
    }