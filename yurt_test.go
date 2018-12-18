package yurt

import (
    "testing"
    "fmt"
)


func TestArgParser(t *testing.T) {
    var p Parser
    called := 0

    p.Command("", "Default", func(_ []string) {
        fmt.Println("Called Default")
        called++
    })

    p.Command("foo", "Foo", func(_ []string) {
        fmt.Println("Called foo")
        called++
    })

    p.Command("foo $s", "Variable Foo", func(_ []string) {
        fmt.Println("Called var foo")
        called++
    })

    p.Run([]string{"a"})
    p.Run([]string{"a", "foo"})
    p.Run([]string{"a", "foo", "bar"})

    if called != 3 {
        t.Fail()
    }
}
