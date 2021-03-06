====
Yurt
====

.. image:: https://img.shields.io/travis/StompRocket/yurt.svg?style=for-the-badge
    :target: https://travis-ci.org/StompRocket/yurt

.. image:: https://img.shields.io/badge/godoc-reference-375EB1.svg?style=for-the-badge
    :target: https://godoc.org/github.com/StompRocket/yurt


A WIP Go CLI library

Docs
----

Read the docs on godoc, click the shield above, or read them offline with
``godoc github.com/StompRocket/yurt`` or ``godoc .`` if you've cloned the repository.

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

Test
----

::

    $ go test
    ╭──────────────╮
    │ Hello, World │
    ╰──────────────╯
    ╭─────╮
    │ Red │
    ╰─────╯
    Called Default
    Called foo
    Called var foo
    PASS