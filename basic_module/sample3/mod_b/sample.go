package mod_b

import (
    // ドットをつけてインポートすると、fmt.PrintではなくPrintとして呼び出せます
    . "fmt"
)

func Test() string {
    // ドットをつけてimportしているので、fmt.がprefixに付与されていないことに注意
    Print("Hello world from mod_b!\n")
    return "This is mod_b Test."
}
