package main

import (
    "fmt"
)

func trace(name string) func() {
    fmt.Printf("%s entered\n", name)
    return func() {
        fmt.Printf("%s returned\n", name)
    }

}
